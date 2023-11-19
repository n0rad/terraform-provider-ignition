package ignition

import (
	"fmt"
	"testing"

	"github.com/coreos/ignition/v2/config/v3_4/types"
)

func TestIgnitionLuks(t *testing.T) {
	testIgnition(t, `
		data "ignition_luks" "foo" {
			name = "foo"
			device = "/foo"
			label = "FOO"
			open_options = ["there"]
			options = ["aes"]
			uuid = "uuid"
			wipe_volume = true
		}

		data "ignition_config" "test" {
			luks = [data.ignition_luks.foo.rendered]
		}
	`, func(c *types.Config) error {
		if len(c.Storage.Luks) != 1 {
			return fmt.Errorf("luks, found %d", len(c.Storage.Luks))
		}

		a := c.Storage.Luks[0]
		if a.Name != "foo" {
			return fmt.Errorf("name, found %q", a.Name)
		}

		if *a.Device != "/foo" {
			return fmt.Errorf("device, found %q", *a.Device)
		}

		if *a.Label != "FOO" {
			return fmt.Errorf("label, found %q", *a.Label)
		}

		if *a.UUID != "uuid" {
			return fmt.Errorf("uuid, found %q", *a.UUID)
		}

		if !*a.WipeVolume {
			return fmt.Errorf("wipeVolume, found %t", *a.WipeVolume)
		}

		if a.OpenOptions[0] != "there" {
			return fmt.Errorf("open_options, found %q", a.OpenOptions)
		}

		if a.Options[0] != "aes" {
			return fmt.Errorf("options, found %q", a.Options)
		}

		return nil
	})
}
