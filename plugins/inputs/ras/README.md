# RAS Daemon Input Plugin

This plugin gathers statistics and error counts provided by the local
[RAS (reliability, availability and serviceability)][ras] daemon.

> [!NOTE]
> This plugin requires access to SQLite3 database from `RASDaemon`. Please make
> sure the Telegraf user has the required permissions to this database!

⭐ Telegraf v1.16.0
🏷️ server
💻 linux

[ras]: https://github.com/mchehab/rasdaemon

## Global configuration options <!-- @/docs/includes/plugin_config.md -->

In addition to the plugin-specific configuration settings, plugins support
additional global and plugin configuration settings. These settings are used to
modify metrics, tags, and field or create aliases and configure ordering, etc.
See the [CONFIGURATION.md][CONFIGURATION.md] for more details.

[CONFIGURATION.md]: ../../../docs/CONFIGURATION.md#plugins

## Configuration

```toml @sample.conf
# RAS plugin exposes counter metrics for Machine Check Errors provided by RASDaemon (sqlite3 output is required).
# This plugin ONLY supports Linux on 386, amd64, arm, and arm64
[[inputs.ras]]
  ## Optional path to RASDaemon sqlite3 database.
  ## Default: /var/lib/rasdaemon/ras-mc_event.db
  # db_path = ""
```

In addition `RASDaemon` runs, by default, with `--enable-sqlite3` flag. In case
of problems with SQLite3 database please verify this is still a default option.

## Metrics

- ras
  - tags:
    - socket_id
  - fields:
    - memory_read_corrected_errors
    - memory_read_uncorrectable_errors
    - memory_write_corrected_errors
    - memory_write_uncorrectable_errors
    - cache_l0_l1_errors
    - tlb_instruction_errors
    - cache_l2_errors
    - upi_errors
    - processor_base_errors
    - processor_bus_errors
    - internal_timer_errors
    - smm_handler_code_access_violation_errors
    - internal_parity_errors
    - frc_errors
    - external_mce_errors
    - microcode_rom_parity_errors
    - unclassified_mce_errors

Please note that `processor_base_errors` is aggregate counter measuring the
following MCE events:

- internal_timer_errors
- smm_handler_code_access_violation_errors
- internal_parity_errors
- frc_errors
- external_mce_errors
- microcode_rom_parity_errors
- unclassified_mce_errors

## Example Output

```text
ras,host=ubuntu,socket_id=0 external_mce_base_errors=1i,frc_errors=1i,instruction_tlb_errors=5i,internal_parity_errors=1i,internal_timer_errors=1i,l0_and_l1_cache_errors=7i,memory_read_corrected_errors=25i,memory_read_uncorrectable_errors=0i,memory_write_corrected_errors=5i,memory_write_uncorrectable_errors=0i,microcode_rom_parity_errors=1i,processor_base_errors=7i,processor_bus_errors=1i,smm_handler_code_access_violation_errors=1i,unclassified_mce_base_errors=1i 1598867393000000000
ras,host=ubuntu level_2_cache_errors=0i,upi_errors=0i 1598867393000000000
```
