package main

const (
	//sync with com.deepin.daemon.power.schema
	PowerPlanCustom          = 0
	PowerPlanPowerSaver      = 1
	PowerPlanBalanced        = 2
	PowerPlanHighPerformance = 3
)

func (p *Power) setBatteryIdleDelay(delay int32) {
	p.setPropBatteryIdleDelay(delay)

	if p.BatteryPlan == PowerPlanCustom && int32(p.coreSettings.GetInt("battery-idle-delay")) != delay {
		p.coreSettings.SetInt("battery-idle-delay", int(delay))
	}
}

func (p *Power) setBatterySuspendDelay(delay int32) {
	p.setPropBatterySuspendDelay(delay)

	if p.BatteryPlan == PowerPlanCustom && int32(p.coreSettings.GetInt("battery-suspend-delay")) != delay {
		p.coreSettings.SetInt("battery-suspend-delay", int(delay))
	}
}

func (p *Power) setBatteryPlan(plan int32) {
	switch plan {
	case PowerPlanHighPerformance:
		p.setBatteryIdleDelay(0)
		p.setBatterySuspendDelay(0)
	case PowerPlanBalanced:
		p.setBatteryIdleDelay(600)
		p.setBatterySuspendDelay(0)
	case PowerPlanPowerSaver:
		p.setBatteryIdleDelay(300)
		p.setBatterySuspendDelay(600)
	case PowerPlanCustom:
		p.setBatteryIdleDelay(int32(p.coreSettings.GetInt("battery-idle-delay")))
		p.setBatterySuspendDelay(int32(p.coreSettings.GetInt("battery-suspend-delay")))
	}
}

func (p *Power) setLinePowerIdleDelay(delay int32) {
	p.setPropLinePowerIdleDelay(delay)

	if p.LinePowerPlan == PowerPlanCustom && int32(p.coreSettings.GetInt("ac-idle-delay")) != delay {
		p.coreSettings.SetInt("ac-idle-delay", int(delay))
	}
}

func (p *Power) setLinePowerSuspendDelay(delay int32) {
	p.setPropLinePowerSuspendDelay(delay)

	if p.LinePowerPlan == PowerPlanCustom && int32(p.coreSettings.GetInt("ac-suspend-delay")) != delay {
		p.coreSettings.SetInt("ac-suspend-delay", int(delay))
	}
}

func (p *Power) setLinePowerPlan(plan int32) {
	switch plan {
	case PowerPlanHighPerformance:
		p.setLinePowerIdleDelay(0)
		p.setLinePowerSuspendDelay(0)
	case PowerPlanBalanced:
		p.setLinePowerIdleDelay(600)
		p.setLinePowerSuspendDelay(0)
	case PowerPlanPowerSaver:
		p.setLinePowerIdleDelay(300)
		p.setLinePowerSuspendDelay(600)
	case PowerPlanCustom:
		p.setLinePowerIdleDelay(int32(p.coreSettings.GetInt("ac-idle-delay")))
		p.setLinePowerSuspendDelay(int32(p.coreSettings.GetInt("ac-suspend-delay")))
	}
}