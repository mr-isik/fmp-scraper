# Hızlı Başlangıç Rehberi

Bu rehber size FMP Scraper'ı 5 dakikada çalıştırmanızı sağlayacak.

## 1️⃣ API Key Alın (2 dakika)

1. [https://site.financialmodelingprep.com/developer/docs](https://site.financialmodelingprep.com/developer/docs) adresine gidin
2. "Get Your Free API Key" butonuna tıklayın
3. Kayıt olun (email gerekli)
4. Dashboard'dan API key'inizi kopyalayın

## 2️⃣ Kurulum (1 dakika)

```bash
# Projeyi indirin
git clone https://github.com/mr-isik/fmp-scraper.git
cd fmp-scraper

# Bağımlılıkları yükleyin
go mod download

# Build edin
go build -o fmp-scraper.exe ./cmd/fmp-scraper
```

## 3️⃣ API Key'i Ayarlayın (30 saniye)

### Yöntem 1: .env Dosyası (Önerilen)

```bash
# .env dosyası oluşturun
cp .env.example .env
```

`.env` dosyasını düzenleyin:

```env
FMP_API_KEY=your_actual_api_key_here
```

### Yöntem 2: PowerShell'de Ortam Değişkeni

```powershell
$env:FMP_API_KEY="your_actual_api_key_here"
```

## 4️⃣ İlk Verilerinizi Çekin! (30 saniye)

```bash
./fmp-scraper.exe -s AAPL -f 2024-01-01 -t 2024-01-31
```

✅ Tebrikler! `AAPL_2024-01-01_2024-01-31.csv` dosyası oluşturuldu.

## 📊 Çıktıyı Kontrol Edin

Excel, Google Sheets veya herhangi bir metin editörü ile CSV dosyasını açın.

## 🎯 Daha Fazla Örnek

### Farklı hisse

```bash
./fmp-scraper.exe -s TSLA -f 2024-01-01 -t 2024-03-31
```

### Özel dosya adı

```bash
./fmp-scraper.exe -s MSFT -f 2023-01-01 -t 2023-12-31 -o microsoft_2023.csv
```

### Son 90 gün

```bash
./fmp-scraper.exe -s GOOGL -f 2024-08-01 -t 2024-11-01 -o google_q3_2024.csv
```

## ❓ Sorun mu Yaşıyorsunuz?

### Hata: "FMP_API_KEY environment variable is required"

✅ **Çözüm:** `.env` dosyasını oluşturdunuz mu? API key'inizi doğru yazdınız mı?

### Hata: "API returned status 401"

✅ **Çözüm:** API key'iniz geçersiz. [FMP Dashboard](https://site.financialmodelingprep.com/developer/docs)'dan yeni bir key alın.

### Hata: "No data found for symbol..."

✅ **Çözüm:** Hisse sembolü doğru mu? (Örn: `AAPL` doğru, `Apple` yanlış)

### Hata: "Invalid date format"

✅ **Çözüm:** Tarih formatı `YYYY-MM-DD` olmalı. Örnek: `2024-01-01`

## 📚 Daha Fazla Bilgi

Detaylı kullanım için [README.md](README.md) dosyasını inceleyin.

---

**İyi kullanımlar! 🚀**
