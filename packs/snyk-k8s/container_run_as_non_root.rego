package main

import data.kubernetes

is_run_as_non_root_disabled[container] {
	not kubernetes.run_as_non_root_at_podlevel
	kubernetes.containers[container]
	not container.securityContext.runAsNonRoot = true
}

is_run_as_non_root_disabled[container] {
	kubernetes.run_as_non_root_at_podlevel
	kubernetes.containers[container]
	container.securityContext.runAsNonRoot == false
}

deny[msg] {
	kubernetes.is_kind_using_template_spec

	is_run_as_non_root_disabled[container]

	msg = sprintf("input.spec.template.spec.containers[%s].securityContext.runAsNonRoot", [container.name])
}

deny[msg] {
	kubernetes.is_kind("Pod")

	is_run_as_non_root_disabled[container]

	msg = sprintf("input.spec.containers[%s].securityContext.runAsNonRoot", [container.name])
}

deny[msg] {
	kubernetes.is_kind("CronJob")

	is_run_as_non_root_disabled[container]

	msg = sprintf("input.spec.jobTemplate.spec.template.spec.containers[%s].securityContext.runAsNonRoot", [container.name])
}