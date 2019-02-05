COMPILE_COMMAND = go build -o bin/cross_compiling 

clean:
	rm -Rf ./bin

# Cross compiling works with env variables and produces a binary for this specific architecture
build-all-binaries: clean
	# doesn't work on my machine and not in travis, see: https://github.com/golang/go/wiki/GoArm
	# GOOS=android GOARCH=arm $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-android-arm
	# GOOS=darwin  GOARCH=arm $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-darwin-arm
	# GOOS=darwin  GOARCH=arm64 $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-darwin-arm64
	GOOS=darwin    GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-darwin-386
	GOOS=darwin    GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-darwin-amd64
	GOOS=dragonfly GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-dragonfly-amd64
	GOOS=freebsd   GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-freebsd-386
	GOOS=freebsd   GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-freebsd-amd64
	GOOS=freebsd   GOARCH=arm      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-freebsd-arm
	GOOS=linux     GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-386
	GOOS=linux     GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-amd64
	GOOS=linux     GOARCH=arm      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-arm
	GOOS=linux     GOARCH=arm64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-arm64
	GOOS=linux     GOARCH=ppc64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-ppc64
	GOOS=linux     GOARCH=ppc64le  $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-ppc64le
	GOOS=linux     GOARCH=mips     $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-mips
	GOOS=linux     GOARCH=mipsle   $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-mipsle
	GOOS=linux     GOARCH=mips64   $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-mips64
	GOOS=linux     GOARCH=mips64le $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-linux-mips64le
	GOOS=netbsd    GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-netbsd-386
	GOOS=netbsd    GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-netbsd-amd64
	GOOS=netbsd    GOARCH=arm      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-netbsd-arm
	GOOS=openbsd   GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-openbsd-386
	GOOS=openbsd   GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-openbsd-amd64
	GOOS=openbsd   GOARCH=arm      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-openbsd-arm
	GOOS=plan9     GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-plan9-386
	GOOS=plan9     GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-plan9-amd64
	GOOS=solaris   GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-solaris-amd64
	GOOS=windows   GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-windows-386
	GOOS=windows   GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/cross_compiling ./bin/cross_compiling-windows-amd64
