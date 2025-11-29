$ErrorActionPreference = 'Stop'

$Repo = "treasureuzoma/readmit"
$Binary = "readmit"
$Os = "windows"
$Arch = "amd64" # Assuming 64-bit for now, can add detection if needed

$Url = "https://github.com/$Repo/releases/latest/download/$Binary-$Os-$Arch.exe"
$InstallDir = "$env:LOCALAPPDATA\readmit"
$ExePath = "$InstallDir\$Binary.exe"

Write-Host "Downloading $Binary..."
if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
}

Invoke-WebRequest -Uri $Url -OutFile $ExePath

# Add to PATH if not present
$UserPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::User)
if ($UserPath -notlike "*$InstallDir*") {
    Write-Host "Adding $InstallDir to PATH..."
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", [EnvironmentVariableTarget]::User)
    $env:Path += ";$InstallDir"
    Write-Host "Please restart your terminal to use the '$Binary' command."
}

Write-Host "Successfully installed $Binary!"
