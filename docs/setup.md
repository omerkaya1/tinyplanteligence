# Setup

## Minimal setup for IDE

### VSCode

```
{
    "go.toolsEnvVars": {
        "GOOS": "linux",
        "GOARCH": "arm",
        "GOROOT": "/Users/user/Library/Caches/tinygo/goroot-f2facaa8251a0546f2afa579c2301669f5df8900338cae36a044442522971ea2",
        "GOFLAGS": "-tags=avr,baremetal,linux,arm,atmega328p,atmega,avr5,arduino,tinygo,purego,osusergo,math_big_pure_go,gc.conservative,scheduler.none,serial.uart"
    }
}
```

## Step-by-step guides

### Setup (Mac)

1. `brew tap tinygo-org/tools`
2. `brew install tinygo`
3. `brew tap osx-cross/avr`
4. `brew install avr-gcc avrdude`

### Flushing

1. Run `tinygo info arduino` to find the information on GOROOT
2. Run `tinygo flash --target=arduino ./cmd/trial/main.go`

### Serial port

#### First use

1. Install Putty - `sudo apt-get install putty` || `brew install Putty`
2. Run Putty in terminal and connect to the serial port: `putty` -> Session -> Connection type -> Serial
3. Locate serial device (arduino - `ls /dev/cu.*`)
4. Save the current or create a new session for re-use.
5. Press `Open`

#### Re-use

1. Run `putty` and select the previously stored session
2. Load the stored session
3. Press `Open`