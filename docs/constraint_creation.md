# How Constraints are Created

## Policy File Parsing

When using `konstraint generate`, Konstraint will only generate templates and constraints for policy files with at least one `violation[]` rule. The `violation` rule is required by Gatekeeper when evaluating policies on a Kubernetes cluster.

When using `konstraint doc`, Konstraint will create documentation for each policy file and assign a severity based on the rule names found in the policy.

The following rule names are organized into their own sections as they have special meaning within the context of Gatekeeper and Conftest:

- `violation` (supported by Gatekeeper and Conftest)
- `warn` (supported by Conftest)

If a policy file does not contain any of the above rules, the policy is added to the `Other` section.

## Importing Libraries

The Rego for the libraries will be added to the generated `ConstraintTemplate` if and only if the policy imports the library. This helps prevent importing Rego code that will go unused.

## Resource Naming

The name of the templates and constraints are derived from the name of the folder that the policy was found in.

For example, a policy found in: `policies/pod-volume-size-limits/src.rego` generates the following in the `policies/pod-volume-size-limits` directory:

- `template.yaml` (defining a ConstraintTemplate)
  - kind: _ConstraintTemplate_
  - name: _podvolumesizelimits_
  - CRD kind (to add to Kubernetes API): _PodVolumeSizeLimits_

- `constraint.yaml` (implementing the above ConstraintTemplate)
  - kind: _PodVolumeSizeLimits_
  - name: _podvolumesizelimits_

When using the `--output` parameter, all templates and constraints will be generated in the path specified in the parameter with the format:

- constraint_PodVolumeSizeLimits.yaml
- template_PodVolumeSizeLimits.yaml

_NOTE: While not technically required, the tool works best with a folder structure similar to how Gatekeeper itself [structures policies and templates](https://github.com/open-policy-agent/gatekeeper/tree/master/library)._

## Annotating Rules

To further promote that the `.rego` file is the source of truth for policy, a block comment can be added to each policy file.

This comment block should:

- Include a human readable description of what the policy does.
- Set the matchers used when generating the Constraints.

It may also specify the enforcment action (either `deny` or `dryrun`) that Gatekeeper should take when a resource violates the constraint. If no enforcement action is specified, Konstraint defaults to using `deny` to align with Gatekeeper's default action. If the enforcement is set to `dryrun`, the policy will be skipped in the documentation generation.

```rego
# @title Pods must not run with access to the host IPC
#
# Pods that are allowed to access the host IPC can read memory of
# the other containers, breaking that security boundary.
#
# @enforcement deny
# @kinds apps/DaemonSet apps/Deployment apps/StatefulSet core/Pod
package pod_deny_host_ipc

import data.lib.core
import data.lib.pods

violation[msg] {
    pod_has_hostipc

    msg := core.format(sprintf("%s/%s: Pod allows for accessing the host IPC", [core.kind, core.name]))
}

pod_has_hostipc {
    pods.pod.spec.hostIPC
}
```

The comment block is also what is used when generating documentation via the `doc` command.
