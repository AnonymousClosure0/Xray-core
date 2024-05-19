
#export CGO_CFLAGS="-fembed-bitcode"
# export GOARCH=amd64
#export SDK=iphonesimulator
#export CARCH="x86_64"  # if compiling for iPhone simulator

# export GOOS=ios
# export SDK=iphoneos

export CGO_ENABLED=1
export GOARCH=arm64

# for ios
export GOOS=ios
export SDK=iphoneos
export IPHONEOS_DEPLOYMENT_TARGET=11.0

# for macos
# export GOOS=darwin
# export SDK=macosx

export SDK_PATH=`xcrun --sdk $SDK --show-sdk-path`
export CLANG=`xcrun --sdk $SDK --find clang`
export CGO_CPPFLAGS="-Wno-error -Wno-nullability-completeness -Wno-expansion-to-defined"
export CGO_CPPFLAGS="-isysroot $SDK_PATH"

go build -buildmode c-archive -o xray ./main
cp xray ./xray.framework/xray
cp xray.h ./xray.framework/Headers/xray.h
cp -r ./xray.framework ../cxray
# cp -r ./xray.framework /Users/jiang/TestXRay/TestXRay