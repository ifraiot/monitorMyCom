# Monitor Your Computer


## Installation
1. Register IFRA IIoT Platform https://app.ifra.io/ and create your asset and measurement with below information.
- Create Asset name "My Computer"
- Create Device name "script"
- Create Measurements (sensor type)
  - memory_total
  - memory_used
  - memory_cached
  - memory_free
  - cpu_sys
  - cpu_user
  - cpu_usage
  - cpu_idle
  - battery_state
  - battery_capacity
  - battery_last_capacity
  - battery_charge_rate


3. Install via go package
```
go install github.com/ifraiot/monitorMyCom
```

4. Start command to get monitoring values.
For topic, username, password, you can copy from IFRA Platform after you create device was completed.

Command Format
```
monitorMyCom -topic=<your-topic> -username=<your-device-username> -password=<your-device-password>
```
  
Example Command
```
monitorMyCom -topic=organization/0693b475-7fed-4749-b8f2-39ead20066c8/messages -username=041123a9-433b-44da-98b5-f9095f28b1ec -password=d581534a-b12a-4c2a-8ad2-d5edf39cf7d4
```

