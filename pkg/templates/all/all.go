package all

import (
	// Import all check templates.
	_ "golang.stackrox.io/kube-linter/pkg/templates/accesstoresources"
	_ "golang.stackrox.io/kube-linter/pkg/templates/antiaffinity"
	_ "golang.stackrox.io/kube-linter/pkg/templates/clusteradminrolebinding"
	_ "golang.stackrox.io/kube-linter/pkg/templates/containercapabilities"
	_ "golang.stackrox.io/kube-linter/pkg/templates/cpurequirements"
	_ "golang.stackrox.io/kube-linter/pkg/templates/danglinghpa"
	_ "golang.stackrox.io/kube-linter/pkg/templates/danglingingress"
	_ "golang.stackrox.io/kube-linter/pkg/templates/danglingnetworkpolicy"
	_ "golang.stackrox.io/kube-linter/pkg/templates/danglingnetworkpolicypeer"
	_ "golang.stackrox.io/kube-linter/pkg/templates/danglingservice"
	_ "golang.stackrox.io/kube-linter/pkg/templates/danglingservicemonitor"
	_ "golang.stackrox.io/kube-linter/pkg/templates/deprecatedserviceaccount"
	_ "golang.stackrox.io/kube-linter/pkg/templates/disallowedgvk"
	_ "golang.stackrox.io/kube-linter/pkg/templates/dnsconfigoptions"
	_ "golang.stackrox.io/kube-linter/pkg/templates/duplicatenvvar"
	_ "golang.stackrox.io/kube-linter/pkg/templates/envvar"
	_ "golang.stackrox.io/kube-linter/pkg/templates/envvarvaluefrom"
	_ "golang.stackrox.io/kube-linter/pkg/templates/forbiddenannotation"
	_ "golang.stackrox.io/kube-linter/pkg/templates/hostipc"
	_ "golang.stackrox.io/kube-linter/pkg/templates/hostmounts"
	_ "golang.stackrox.io/kube-linter/pkg/templates/hostnetwork"
	_ "golang.stackrox.io/kube-linter/pkg/templates/hostpid"
	_ "golang.stackrox.io/kube-linter/pkg/templates/hpareplicas"
	_ "golang.stackrox.io/kube-linter/pkg/templates/imagepullpolicy"
	_ "golang.stackrox.io/kube-linter/pkg/templates/latesttag"
	_ "golang.stackrox.io/kube-linter/pkg/templates/livenessport"
	_ "golang.stackrox.io/kube-linter/pkg/templates/livenessprobe"
	_ "golang.stackrox.io/kube-linter/pkg/templates/memoryrequirements"
	_ "golang.stackrox.io/kube-linter/pkg/templates/mismatchingselector"
	_ "golang.stackrox.io/kube-linter/pkg/templates/namespace"
	_ "golang.stackrox.io/kube-linter/pkg/templates/nodeaffinity"
	_ "golang.stackrox.io/kube-linter/pkg/templates/nonexistentserviceaccount"
	_ "golang.stackrox.io/kube-linter/pkg/templates/nonisolatedpod"
	_ "golang.stackrox.io/kube-linter/pkg/templates/pdbmaxunavailable"
	_ "golang.stackrox.io/kube-linter/pkg/templates/pdbminavailable"
	_ "golang.stackrox.io/kube-linter/pkg/templates/ports"
	_ "golang.stackrox.io/kube-linter/pkg/templates/privileged"
	_ "golang.stackrox.io/kube-linter/pkg/templates/privilegedports"
	_ "golang.stackrox.io/kube-linter/pkg/templates/privilegeescalation"
	_ "golang.stackrox.io/kube-linter/pkg/templates/readinessport"
	_ "golang.stackrox.io/kube-linter/pkg/templates/readinessprobe"
	_ "golang.stackrox.io/kube-linter/pkg/templates/readonlyrootfs"
	_ "golang.stackrox.io/kube-linter/pkg/templates/readsecret"
	_ "golang.stackrox.io/kube-linter/pkg/templates/replicas"
	_ "golang.stackrox.io/kube-linter/pkg/templates/requiredannotation"
	_ "golang.stackrox.io/kube-linter/pkg/templates/requiredlabel"
	_ "golang.stackrox.io/kube-linter/pkg/templates/runasnonroot"
	_ "golang.stackrox.io/kube-linter/pkg/templates/sccdenypriv"
	_ "golang.stackrox.io/kube-linter/pkg/templates/serviceaccount"
	_ "golang.stackrox.io/kube-linter/pkg/templates/servicetype"
	_ "golang.stackrox.io/kube-linter/pkg/templates/startupport"
	_ "golang.stackrox.io/kube-linter/pkg/templates/sysctl"
	_ "golang.stackrox.io/kube-linter/pkg/templates/targetport"
	_ "golang.stackrox.io/kube-linter/pkg/templates/unsafeprocmount"
	_ "golang.stackrox.io/kube-linter/pkg/templates/updateconfig"
	_ "golang.stackrox.io/kube-linter/pkg/templates/wildcardinrules"
	_ "golang.stackrox.io/kube-linter/pkg/templates/writablehostmount"
)
