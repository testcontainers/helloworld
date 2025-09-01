ARG BASE_IMAGE

# -----------
# Build Image
# -----------
FROM golang:1.25-nanoserver as builder
ENV CGO_ENABLED=0

SHELL ["powershell", "-Command", "$ErrorActionPreference = 'Stop'; $ProgressPreference = 'SilentlyContinue';"]

WORKDIR /app

# Install source deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source & build
COPY . .
# Build binary:
# -a: force rebuild
# -installsuffix cgo: remove cgo support
# -ldflags="-w -s -extldflags '-static'": enable static linking, omits the DWARF symbol table
# -trimpath: remove all file system paths from the compiled executable
# -o /bin/helloworld: output binary to /bin/helloworld
RUN go build \
    -a \
    -installsuffix cgo \
    -ldflags="-w -s -extldflags '-static'" \
    -trimpath \
    -o /bin/helloworld .

# Try to compress with UPX, but don't fail if it doesn't work
RUN try { \
        Write-Host 'Attempting to download and use UPX...'; \
        Invoke-WebRequest -Uri 'https://github.com/upx/upx/releases/download/v5.0.2/upx-5.0.2-win64.zip' -OutFile 'upx.zip' -UseBasicParsing; \
        Expand-Archive -Path 'upx.zip' -DestinationPath 'C:\upx' -Force; \
        $upxExe = Get-ChildItem -Path 'C:\upx' -Name 'upx.exe' -Recurse | Select-Object -First 1; \
        if ($upxExe) { \
            $upxPath = Join-Path 'C:\upx' $upxExe; \
            & $upxPath --best --lzma /bin/helloworld; \
            Write-Host 'UPX compression successful'; \
        } else { \
            Write-Host 'UPX executable not found, skipping compression'; \
        } \
    } catch { \
        Write-Host "UPX compression failed: $_. Continuing without compression."; \
    } finally { \
        Remove-Item -Path 'upx.zip' -Force -ErrorAction SilentlyContinue; \
        Remove-Item -Path 'C:\upx' -Recurse -Force -ErrorAction SilentlyContinue; \
    }

# -----------------
# Distributed Image
# -----------------
FROM ${BASE_IMAGE}
COPY static /static
COPY --from=builder /bin/helloworld /bin/helloworld
ENTRYPOINT ["/bin/helloworld"]
