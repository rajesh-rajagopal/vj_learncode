EMBEDDED_CONFIG=y;DIND_IMAGE=mirantis/kubeadm-dind-cluster:v1.6

CNI_PLUGIN="${CNI_PLUGIN:-bridge}"
DIND_SUBNET="${DIND_SUBNET:-10.192.0.0}"
dind_ip_base="$(echo "${DIND_SUBNET}" | sed 's/\.0$//')"
DIND_IMAGE="${DIND_IMAGE:-}"
BUILD_KUBEADM="${BUILD_KUBEADM:-}"
BUILD_HYPERKUBE="${BUILD_HYPERKUBE:-}"
APISERVER_PORT=${APISERVER_PORT:-8080}
NUM_NODES=${NUM_NODES:-2}
LOCAL_KUBECTL_VERSION=${LOCAL_KUBECTL_VERSION:-}
KUBECTL_DIR="${KUBECTL_DIR:-${HOME}/.kubeadm-dind-cluster}"
DASHBOARD_URL="${DASHBOARD_URL:-https://rawgit.com/kubernetes/dashboard/bfab10151f012d1acc5dfb1979f3172e2400aa3c/src/deploy/kubernetes-dashboard.yaml}"
E2E_REPORT_DIR="${E2E_REPORT_DIR:-}"


function dind::check-binary {
  local filename="$1"
  local dockerized="_output/dockerized/bin/linux/amd64/${filename}"
  local plain="_output/local/bin/linux/amd64/${filename}"
  dind::set-build-volume-args
  # FIXME: don't hardcode amd64 arch
  if [ -n "${KUBEADM_DIND_LOCAL:-${force_local:-}}" ]; then
    echo "yes kubeadm dind local"
    if [ -f "${dockerized}" -o -f "${plain}" ]; then
      return 0
    fi
  elif docker run --rm "${build_volume_args[@]}" \
              "${busybox_image}" \
              test -f "/go/src/k8s.io/kubernetes/${dockerized}" >&/dev/null; then
    return 0
  fi
  return 1
}

function dind::volume-exists {
  local name="$1"
  if docker volume inspect "${name}" >& /dev/null; then
   echo "* volome if"
    return 0
  fi
  echo "** volume exists ${name}"
  return 1
}


if ! dind::volume-exists kubeadm-dind-sys; then
   echo "sdfssdfd"
fi
