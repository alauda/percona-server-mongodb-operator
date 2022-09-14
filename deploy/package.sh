#######
# Script for packaging mongo related images together and upload into S3 object storage
#######
# Unit Test section: prepare some fake data
#######
mkdir -p bundle/manifests/
cat > bundle/manifests/percona-server-mongodb-operator.clusterserviceversion.yaml << EOF
spec:
    relatedImages:
    - name: values-yaml-mongo-backup
      image: build-harbor.alauda.cn/middleware/percona-backup-mongodb:v1.3.4-mongodb-backup
      support_arm: true
    - name: values-yaml-mongo-exporter
      image: build-harbor.alauda.cn/middleware/mongodb-exporter:v0.20.1-mongodb-exporter
      support_arm: true
    - name: values-yaml-mongo-operator
      image: build-harbor.alauda.cn/middleware/percona-server-mongodb-operator:v3.10-15-g221feea6
      support_arm: true
    - name: values-yaml-mongo-server-4.2
      image: build-harbor.alauda.cn/middleware/percona-server-mongodb:v4.2.17-mongodb
      support_arm: true
    - name: values-yaml-mongo-server-4.4
      image: build-harbor.alauda.cn/middleware/percona-server-mongodb:v4.4.6-mongodb
      support_arm: true
EOF

#######
# Package logic below which is executed in pipeline
#######
set -ex
export BACKUP_AGENT_IMAGE=`yq r bundle/manifests/percona-server-mongodb-operator.clusterserviceversion.yaml 'spec.relatedImages[0].image' `
export EXPORTER_IMAGE=`yq r bundle/manifests/percona-server-mongodb-operator.clusterserviceversion.yaml 'spec.relatedImages[1].image' `
export OPERATOR_IMAGE=`yq r bundle/manifests/percona-server-mongodb-operator.clusterserviceversion.yaml 'spec.relatedImages[2].image' `
export MONGO_SERVER_42_IMAGE=`yq r bundle/manifests/percona-server-mongodb-operator.clusterserviceversion.yaml 'spec.relatedImages[3].image' `
export MONGO_SERVER_44_IMAGE=`yq r bundle/manifests/percona-server-mongodb-operator.clusterserviceversion.yaml 'spec.relatedImages[4].image' `
export LOCAL_BACKUP_PKG_FILE=`echo \$BACKUP_AGENT_IMAGE | awk -F "/" '{print \$3}' | tr ":" "."`
echo -e "backup image: \$LOCAL_BACKUP_PKG_FILE \n"
export LOCAL_EXPORTER_PKG_FILE=`echo \$EXPORTER_IMAGE | awk -F "/" '{print \$3}' | tr ":" "."`
echo -e "exporter image: \$LOCAL_EXPORTER_PKG_FILE \n"
export LOCAL_OPERATOR_PKG_FILE=`echo \$OPERATOR_IMAGE | awk -F "/" '{print \$3}' | tr ":" "."`
echo -e "operator image: \$LOCAL_OPERATOR_PKG_FILE \n"
export LOCAL_BUNDLE_VERSION=`echo \$OPERATOR_IMAGE | awk -F "/" '{print \$3}' | awk -F ":" '{print \$2}'`
echo -e "local bundle version \$LOCAL_BUNDLE_VERSION"
export LOCAL_MONGO_SVR42_FILE=`echo \$MONGO_SERVER_42_IMAGE | awk -F "/" '{print \$3}' | tr ":" "."`
echo -e "mongo server 4.2 image: \$LOCAL_MONGO_SVR42_FILE \n"
export LOCAL_MONGO_SVR44_FILE=`echo \$MONGO_SERVER_44_IMAGE | awk -F "/" '{print \$3}' | tr ":" "."`
echo -e "mongo server 4.4 image: \$LOCAL_MONGO_SVR44_FILE \n"
export LOCAL_BUNDLE_IMAGE="percona-server-mongodb-operator-bundle.\${LOCAL_BUNDLE_VERSION}"
echo -e "local bundle image: \${LOCAL_BUNDLE_IMAGE} \n"
export BUNDLE_IMAGE="build-harbor.alauda.cn/middleware/percona-server-mongodb-operator-bundle:\${LOCAL_BUNDLE_VERSION}"
echo -e "bundle image: \${BUNDLE_IMAGE} \n"

export OUTDIR="\${LOCAL_BUNDLE_VERSION}"
mkdir -p \${OUTDIR}
export IMAGE_CMD=skopeo
\$IMAGE_CMD copy -a docker://\${BACKUP_AGENT_IMAGE} dir:\${OUTDIR}/\${LOCAL_BACKUP_PKG_FILE}
\$IMAGE_CMD copy -a docker://\${EXPORTER_IMAGE} dir:\${OUTDIR}/\${LOCAL_EXPORTER_PKG_FILE}
\$IMAGE_CMD copy -a docker://\${OPERATOR_IMAGE} dir:\${OUTDIR}/\${LOCAL_OPERATOR_PKG_FILE}
\$IMAGE_CMD copy -a docker://\${MONGO_SERVER_42_IMAGE} dir:\${OUTDIR}/\${LOCAL_MONGO_SVR42_FILE}
\$IMAGE_CMD copy -a docker://\${MONGO_SERVER_44_IMAGE} dir:\${OUTDIR}/\${LOCAL_MONGO_SVR44_FILE}
\$IMAGE_CMD copy -a docker://\${BUNDLE_IMAGE} dir:\${OUTDIR}/\${LOCAL_BUNDLE_IMAGE}

echo -e "Begin packaging...\n"
tar zcvf \${OUTDIR}".tar" ./\${OUTDIR}
echo -e "Prepare upload tool mc...\n"
curl https://dl.minio.io/client/mc/release/linux-amd64/mc -o mc
chmod +x mc
echo -e "package \${OUTDIR}.tar done, start uploading...\n"
# Remember to update below key&secret before use
./mc config host add idc http://192.168.144.3:28001 S3KEY S3SECRET --api s3v2
./mc cp \${OUTDIR}".tar" idc/3.0/ 
echo -e "Upload finished succesfully!\n"

