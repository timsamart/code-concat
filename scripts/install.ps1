# Determine the latest version (you would need to implement this based on your release strategy)
$VERSION = "v1.0.0"

# Determine system architecture
$ARCH = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

# Set the download URL
$URL = "https://github.com/timsamart/code-concat/releases/download/${VERSION}/code-concat_windows_${ARCH}.exe"

# Set the installation directory
$INSTALL_DIR = "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps"

# Download the binary
Write-Host "Downloading Code-Concat..."
Invoke-WebRequest -Uri $URL -OutFile "codeconcat.exe"

# Move the binary to the installation directory
Move-Item -Path "codeconcat.exe" -Destination "$INSTALL_DIR\codeconcat.exe" -Force

Write-Host "Code-Concat has been installed to $INSTALL_DIR\codeconcat.exe"
Write-Host "You can now use it by running 'codeconcat' in your PowerShell or Command Prompt."

# Add the installation directory to the PATH if it's not already there
if ($env:Path -notlike "*$INSTALL_DIR*") {
    [Environment]::SetEnvironmentVariable("Path", $env:Path + ";$INSTALL_DIR", [EnvironmentVariableTarget]::User)
    Write-Host "Added $INSTALL_DIR to your PATH. You may need to restart your terminal for this change to take effect."
}