package datadog

import (
	"fmt"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/processors"
)

const sampleConfig = `
`

type Datadog struct {
}

func (r *Datadog) SampleConfig() string {
	return sampleConfig
}

func (r *Datadog) Description() string {
	return "Rename host measurements with datadog equivalent."
}

func renameField(metric telegraf.Metric, before string, after string) {
	if value, ok := metric.GetField(before); ok {
		fmt.Printf("Renaming '%s'...\n", before)
		metric.RemoveField(before)
		metric.AddField(after, value)
	}
}

func renameTag(metric telegraf.Metric, before string, after string) {
	if value, ok := metric.GetTag(before); ok {
		metric.RemoveTag(before)
		metric.AddTag(after, value)
	}
}


func (r *Datadog) Apply(in ...telegraf.Metric) []telegraf.Metric {
	//fmt.Printf("\n\n\n\nDATADOG PROCESSOR\n")
	for _, point := range in {
		//fmt.Printf("name... %s\n", point.Name())
		switch point.Name() {
		case "system":
			renameField(point, "load1", "load.1")
			renameField(point, "load5", "load.5")
			renameField(point, "load15", "load.15")
		case "cpu":
			if _, ok := point.GetTag("cpu"); ok {
				point.SetName("system")
				renameField(point, "usage_idle", "cpu.idle")
				renameField(point, "usage_system", "cpu.system")
				renameField(point, "usage_iowait", "cpu.iowait")
				renameField(point, "usage_user", "cpu.user")
				renameField(point, "usage_steal", "cpu.stolen")
				renameField(point, "usage_guest", "cpu.guest")
				renameField(point, "usage_guest_nice", "cpu.guest.nice")
				renameField(point, "usage_nice", "cpu.nice")
				renameField(point, "usage_softirq", "cpu.softirq")
				renameField(point, "usage_irq", "cpu.irq")
			}
		case "mem":
			if _, ok := point.GetField("used"); ok {
				point.SetName("system")
				renameField(point, "available", "mem.usable")
				renameField(point, "total", "mem.total")
				renameField(point, "free", "mem.free")
				renameField(point, "used", "mem.used")
				renameField(point, "commit_limit", "mem.commit.limit")
				renameField(point, "shared", "mem.shared")
				renameField(point, "used_percent", "mem.used.percent")
				renameField(point, "mapped", "mem.mapped")
				renameField(point, "active", "mem.active")
				renameField(point, "huge_page_size", "mem.huge_page_size")
				renameField(point, "vmalloc_used", "mem.vmalloc_used")
				renameField(point, "write_back", "mem.write_back")
				renameField(point, "committed_as", "mem.committed_as")
				renameField(point, "dirty", "mem.dirty")
				renameField(point, "page_tables", "mem.page_tables")
				renameField(point, "inactive", "mem.inactive")
				renameField(point, "wired", "mem.wired")
				renameField(point, "slab", "mem.slab")
				renameField(point, "available_percent", "mem.available.percent")
				renameField(point, "huge_pages_total", "mem.huge_pages_total")
				renameField(point, "low_free", "mem.low_free")
				renameField(point, "low_total", "mem.low_total")
				renameField(point, "swap_free", "mem.swap_free")
				renameField(point, "cached", "mem.cached")
				renameField(point, "huge_pages_free", "mem.huge_pages_free")
				renameField(point, "vmalloc_total", "mem.vmalloc_total")
				renameField(point, "write_back", "mem.write_back")
				renameField(point, "swap_cached", "mem.swap_cached")
				renameField(point, "swap_total", "mem.swap_total")
				renameField(point, "vmalloc_chunk", "mem.vmalloc_chunk")
				renameField(point, "buffered", "mem.buffered")
				renameField(point, "high_free", "mem.high_free")
				renameField(point, "high_total", "mem.high_total")
			}
		case "swap":
			if _, ok := point.GetField("free"); ok {
				point.SetName("system")
				renameField(point, "free", "swap.free")
				renameField(point, "used", "swap.used")
				renameField(point, "total", "swap.total")
				renameField(point, "used_percent", "swap.used.percent")
			}
		case "diskio":
			if _, ok := point.GetField("io_time"); ok {
				point.SetName("system")
				renameField(point, "io_time", "io.await")
				renameField(point, "reads", "io.reads")
				renameField(point, "read_bytes", "io.read.bytes")
				renameField(point, "write_bytes", "io.write.bytes")
				renameField(point, "read_time", "io.read.time")
				renameField(point, "writes", "io.writes")
				renameField(point, "write_time", "io.write.time")
				renameField(point, "weighted_io_time", "io.weighted.io.time")
				renameField(point, "iops_in_progress", "io.iops.in.progress")
			}
		case "disk":
			//disk,device=nvme0n1p1,fstype=vfat,host=will-algorand,mode=rw,path=/boot/efi
			if _, ok := point.GetField("used"); ok {
				point.SetName("system")
				renameField(point, "total", "disk.total")
				renameField(point, "free", "disk.free")
				renameField(point, "used", "disk.in_use")
				renameField(point, "used_percent", "disk.used.percent")
				renameField(point, "inodes_total", "disk.inodes.total")
				renameField(point, "inodes_free", "disk.inodes.free")
				renameField(point, "inodes_used", "disk.inodes.used")
			}
		}
	}
	return in
}


func init() {
	processors.Add("datadog", func() telegraf.Processor {
		return &Datadog{}
	})
}
