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

$ZipName = "$Repo-$Version-$OS-$Arch.tar.gz"
$Url = "https://github.com/$Owner/$Repo/releases/download/$Version/$ZipName"

Invoke-WebRequest -Uri $Url -OutFile $ZipName
tar -xzf $ZipName
Move-Item -Force ".\$Repo.exe" "$env:USERPROFILE\.local\bin\$Repo.exe"

$envPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($envPath -notlike "*$env:USERPROFILE\.local\bin*") {
  [Environment]::SetEnvironmentVariable("PATH", "$envPath;$env:USERPROFILE\.local\bin", "User")
}

Write-Output "Installed $Repo $Version to $env:USERPROFILE\.local\bin"
