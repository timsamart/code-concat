# Determine the latest version and system architecture (replace with your logic)
$VERSION = "v1.0.2"
$ARCH = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

# Set the download URL and installation directory
$URL = "https://github.com/timsamart/code-concat/releases/download/${VERSION}/code-concat_windows_${ARCH}.exe"
$INSTALL_DIR = "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps"

# Check if the file exists before downloading again
$filePath = Join-Path -Path $PSScriptRoot -ChildPath "codeconcat.exe"

if (-not (Test-Path $filePath)) {
    try {
        # Download the binary
        Write-Host "Downloading Code-Concat..."
        Invoke-WebRequest -Uri $URL -OutFile $filePath -ErrorAction Stop

        Write-Host "Code-Concat downloaded successfully."
    } catch {
        Write-Host "Error downloading Code-Concat: $_"
    }
} else {
    Write-Host "Code-Concat already exists in $PSScriptRoot."
}

# Move the binary to the installation directory if downloaded successfully
if (Test-Path $filePath) {
    try {
        Move-Item -Path $filePath -Destination "$INSTALL_DIR\codeconcat.exe" -Force
        Write-Host "Code-Concat has been installed to $INSTALL_DIR\codeconcat.exe"
    } catch {
        Write-Host "Error moving Code-Concat: $_"
    }
}

# Optionally, add $INSTALL_DIR to the PATH if needed (requires admin privileges)
# Ensure admin privileges if modifying Path environment variable
if ($env:Path -notlike "*$INSTALL_DIR*") {
    [Environment]::SetEnvironmentVariable("Path", $env:Path + ";$INSTALL_DIR", [EnvironmentVariableTarget]::User)
    Write-Host "Added $INSTALL_DIR to your PATH. You may need to restart your terminal for this change to take effect."
}
