# Toplu veri çekme örneği
# Bu script birden fazla hisse için veri çeker

# API Key kontrolü
if (-not $env:FMP_API_KEY) {
    Write-Host "HATA: FMP_API_KEY ortam degiskeni tanimli degil!" -ForegroundColor Red
    Write-Host "Once .env dosyasini olusturun veya asagidaki komutu calistirin:" -ForegroundColor Yellow
    Write-Host '$env:FMP_API_KEY="your_api_key_here"' -ForegroundColor Cyan
    exit 1
}

# Çekilecek hisse sembolleri
$symbols = @("AAPL", "TSLA", "MSFT", "GOOGL", "AMZN", "META", "NVDA")

# Tarih aralığı
$fromDate = "2024-01-01"
$toDate = "2024-11-15"

# Çıktı klasörü oluştur
$outputDir = "output"
if (-not (Test-Path $outputDir)) {
    New-Item -ItemType Directory -Path $outputDir | Out-Null
    Write-Host "Cikti klasoru olusturuldu: $outputDir" -ForegroundColor Green
}

Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "FMP Scraper - Toplu Veri Cekme" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

Write-Host "Tarih Araligi: $fromDate - $toDate" -ForegroundColor Yellow
Write-Host "Hisse Sayisi: $($symbols.Count)" -ForegroundColor Yellow
Write-Host "Hisseler: $($symbols -join ', ')`n" -ForegroundColor Yellow

$successCount = 0
$errorCount = 0

foreach ($symbol in $symbols) {
    $outputFile = "$outputDir\${symbol}_${fromDate}_${toDate}.csv"
    
    Write-Host "[$($symbols.IndexOf($symbol) + 1)/$($symbols.Count)] $symbol verisi cekiliyor..." -ForegroundColor White -NoNewline
    
    try {
        # FMP Scraper'ı çalıştır
        & .\fmp-scraper.exe -s $symbol -f $fromDate -t $toDate -o $outputFile 2>&1 | Out-Null
        
        if ($LASTEXITCODE -eq 0) {
            Write-Host " BASARILI" -ForegroundColor Green
            $successCount++
        } else {
            Write-Host " HATA" -ForegroundColor Red
            $errorCount++
        }
    }
    catch {
        Write-Host " HATA: $_" -ForegroundColor Red
        $errorCount++
    }
    
    # API rate limiting için kısa bekleme
    Start-Sleep -Milliseconds 500
}

Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "OZET" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Basarili: $successCount" -ForegroundColor Green
Write-Host "Hata: $errorCount" -ForegroundColor Red
Write-Host "Toplam: $($symbols.Count)" -ForegroundColor Yellow
Write-Host "`nCikti klasoru: $outputDir" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

if ($successCount -gt 0) {
    Write-Host "Dosyalar basariyla olusturuldu! Klasoru acmak icin:" -ForegroundColor Green
    Write-Host "explorer $outputDir" -ForegroundColor Yellow
}
