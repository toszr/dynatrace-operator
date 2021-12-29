package statefulset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtensionController_BuildContainerAndVolumes(t *testing.T) {
	assertion := assert.New(t)

	instance := buildTestInstance()
	capabilityProperties := &instance.Spec.Routing.CapabilityProperties
	stsProperties := NewStatefulSetProperties(instance, capabilityProperties,
		"", "", "", "", "",
		nil, nil, nil,
	)

	t.Run("happy path", func(t *testing.T) {
		eec := NewExtensionController(stsProperties)
		container := eec.BuildContainer()

		assertion.NotEmpty(container.ReadinessProbe, "Expected readiness probe is defined")
		assertion.Equal("/readyz", container.ReadinessProbe.Handler.HTTPGet.Path, "Expected there is a readiness probe at /readyz")
		assertion.Empty(container.LivenessProbe, "Expected there is no liveness probe (not implemented)")
		assertion.Empty(container.StartupProbe, "Expected there is no startup probe")

		for _, port := range []int32{eecIngestPort} {
			assertion.Truef(portIsIn(container.Ports, port), "Expected that EEC container defines port %d", port)
		}

		for _, mountPath := range []string{
			activeGateConfigDir,
			"/mnt/dsexecargs",
			"/var/lib/dynatrace/remotepluginmodule/agent/runtime/datasources",
			"/opt/dynatrace/remotepluginmodule/agent/datasources/statsd",
			extensionsLogsDir,
			statsDLogsDir,
		} {
			assertion.Truef(mountPathIsIn(container.VolumeMounts, mountPath), "Expected that EEC container defines mount point %s", mountPath)
		}

		for _, envVar := range []string{
			"TenantId", "ServerUrl", "EecIngestPort",
		} {
			assertion.Truef(envVarIsIn(container.Env, envVar), "Expected that EEC container defined environment variable %s", envVar)
		}
	})

	t.Run("volumes vs volume mounts", func(t *testing.T) {
		eec := NewExtensionController(stsProperties)
		statsd := NewStatsD(stsProperties)
		volumes := buildVolumes(stsProperties, []ContainerBuilder{eec, statsd})

		container := eec.BuildContainer()
		for _, volumeMount := range container.VolumeMounts {
			assertion.Truef(volumeIsDefined(volumes, volumeMount.Name), "Expected that volume mount %s has a predefined pod volume", volumeMount.Name)
		}
	})
}
