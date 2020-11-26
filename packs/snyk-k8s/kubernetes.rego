package kubernetes

pod_containers(pod) = all_containers {
	keys = {"containers", "initContainers"}
	all_containers = [c | keys[k]; c = pod.spec[k][_]]
}

containers[container] {
	pods[pod]
	all_containers = pod_containers(pod)
	container = all_containers[_]
}

containers[container] {
	all_containers = pod_containers(input)
	container = all_containers[_]
}

pods[pod] {
	is_kind_using_template_spec
	pod = input.spec.template
}

pods[pod] {
	is_kind("Pod")
	pod = input
}

pods[pod] {
	is_kind("CronJob")
	pod = input.spec.jobTemplate.spec.template
}

services[service] {
	is_kind("Service")
	service = input
}

volumes[volume] {
	pods[pod]
	volume = pod.spec.volumes[_]
}

added_capability(container, cap) {
	container.securityContext.capabilities.add[_] == cap
}

dropped_capability(container, cap) {
	container.securityContext.capabilities.drop[_] == cap
}

# priviledge_escalation_allowed(c) {
# 	not has_field(c, "securityContext")
# }

priviledge_escalation_allowed(c) {
	has_field(c, "securityContext")
	has_field(c.securityContext, "allowPrivilegeEscalation")
}

no_read_only_filesystem(c) {
	not has_field(c, "securityContext")
}

no_read_only_filesystem(c) {
	has_field(c, "securityContext")
	not has_field(c.securityContext, "readOnlyRootFilesystem")
}

run_as_non_root_at_podlevel() {
	pods[pod]
	has_field(pod.spec, "securityContext")
	pod.spec.securityContext.runAsNonRoot == true
}

is_netpol_api_version_supported {
	supported_api_versions := {"extensions/v1beta1", "networking.k8s.io/v1"}
	some i
	input.apiVersion == supported_api_versions[i]
}

# Check for premissive rules 
## FIXME: other permutations of rules that lead to allow-all
risky_ingres_egress_rules = {{}, [] ,{"from": []}, {"to": []}}
is_egress_block {
  has_field(input.spec, "egress")
}

any_risky_egress_rules {
  some i, j
  input.spec.egress[i] == risky_ingres_egress_rules[j]
}

is_ingress_block {
  has_field(input.spec, "ingress")
}

any_risky_ingress_rules {
  some i, j
  input.spec.ingress[i] == risky_ingres_egress_rules[j]
}

is_psp_api_version_supported {
	supported_api_versions := {"policy/v1beta1"}
	some i
	input.apiVersion == supported_api_versions[i]
}

template_spec_workloads := [
	"Deployment",
	"DaemonSet",
	"StatefulSet",
	"ReplicaSet",
	"ReplicationController",
	"Job"
]

is_kind_using_template_spec {
	some i
	template_spec_workloads[i] == input.kind
}

is_kind(kind) {
    input.kind = kind
}

has_field(obj, field) {
	obj[field]
}

is_field_have_value(obj, field, value) {
	obj[field] = value
}

is_field_in_spec(field, value) {
	has_field(input, "spec")
    has_field(input.spec, field)
    is_field_have_value(input.spec, field, value)
}

is_field_in_metadata_annotations(field) {
    has_field(input, "metadata")
    has_field(input.metadata, "annotations")
    has_field(input.metadata.annotations, field)
}

is_value_in_string(obj, field, value, delimiter) {
    ranges := split(obj[field], delimiter)
  	some i 
    trim_space(ranges[i]) == value
}

is_field_in_spec_with_sub(field, sub_field, value) {
    has_field(input, "spec")
    has_field(input.spec, field)
    has_field(input.spec[field], sub_field)
    is_field_have_value(input.spec[field], sub_field, value)
}

is_field_not_in_spec(field, value) {
	has_field(input, "spec")
    has_field(input.spec, field)
    not is_field_have_value(input.spec, field, value)
}