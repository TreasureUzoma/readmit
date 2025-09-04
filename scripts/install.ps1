param(
  [string]$Version = "latest"
)

$Owner = "TreasureUzoma"
$Repo = "readmit"

Write-Output "📦 Installing $Repo ($Version)..."

if ($Version -eq "latest") {
  $Version = (Invoke-RestMethod "https://api.github.com/repos/$Owner/$Repo/releases/latest").tag_name
}

$Arch = if ($env:PROCESSOR_ARCHITECTURE -eq "AMD64") { "amd64" } else { "arm64" }
$OS = "windows"

$ZipName = "$Repo-$Version-$OS-$Arch.zip"
$Url = "https://github.com/$Owner/$Repo/releases/download/$Version/$ZipName"

$TempDir = New-Item -ItemType Directory -Force -Path ([System.IO.Path]::GetTempPath() + [System.Guid]::NewGuid())
$ZipPath = Join-Path $TempDir $ZipName

Invoke-WebRequest -Uri $Url -OutFile $ZipPath
Expand-Archive -Path $ZipPath -DestinationPath $TempDir -Force

$BinDir = "$env:USERPROFILE\.local\bin"
New-Item -ItemType Directory -Force -Path $BinDir | Out-Null
Move-Item -Force (Join-Path $TempDir "$Repo.exe") (Join-Path $BinDir "$Repo.exe")

$envPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($envPath -notlike "*$BinDir*") {
  [Environment]::SetEnvironmentVariable("PATH", "$envPath;$BinDir", "User")
}

Write-Output "✅ Installed $Repo $Version to $BinDir"
Write-Output "Please restart your terminal or run `& { $env:PATH = [Environment]::GetEnvironmentVariable('PATH', 'User') }` to update your PATH."