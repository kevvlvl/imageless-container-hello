# Creating an imageless container

A special thank you and credits for this example goes to Ivan Velichko and his blog: https://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/

## Steps

1. Generate an oci/config config file
```
    runc spec
```

2. Create a Go file with a simple "Hello world" example. Make sure you have go installed at this point
3. Setup the container's filesystem.
```
$ go mod init
$ GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o printme
$ file printme
> printme: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, stripped

$ mkdir rootfs
$ mv printme rootfs/
```

4. Change the config.json file 
   - process.args = "/printme"
   - hostname = "imageless-hello"
   - process.terminal = false"
5. Create the container
```
sudo runc create imageless-hello
```

6. We can now list the container that is created
```
sudo runc list                 
ID                PID         STATUS      BUNDLE                                 CREATED                          OWNER
imageless-hello   0           created     /home/kev/go/src/imageless-container   2021-11-16T05:03:50.605474896Z   root
```

7. Run the container
```
sudo runc start imageless-hello
```

At this point you should see that the OCI/container we ran outputted "Operating System Hostname imageless-hello"

8. Now we can delete the container that is in stopped STATUS
```
sudo runc delete imageless-hello
```