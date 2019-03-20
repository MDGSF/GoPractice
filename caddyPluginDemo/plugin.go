package caddyPluginDemo

import "github.com/mholt/caddy"

func init() {
	caddy.RegisterPlugin("gizmo", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	for c.Next() {
		if !c.NextArg() {
			return c.ArgErr()
		}
		value := c.Val()
	}
	return nil
}
