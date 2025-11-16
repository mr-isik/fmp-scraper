# FMP Scraper

<div align="center">

📊 **Gelişmiş Financial Modeling Prep API Veri Çekici**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Modern, hızlı ve kullanımı kolay bir CLI tool ile finansal verileri CSV formatında dışa aktarın.

</div>

---

## 🌟 Özellikler

- ✅ **Tarih Aralığı Desteği**: Belirli tarih aralıklarından veri çekme
- 📁 **Özelleştirilebilir Çıktı**: CSV dosya adını dilediğiniz gibi ayarlayın
- 🏗️ **SOLID Prensipler**: Temiz mimari ve sürdürülebilir kod yapısı
- 🚀 **Yüksek Performans**: Go'nun hızından tam anlamıyla yararlanma
- 🔒 **Güvenli**: API anahtarları çevre değişkenleri ile yönetilir
- 📊 **Detaylı Veri**: OHLC, Volume, VWAP ve daha fazlası
- 🎯 **Kullanıcı Dostu**: Basit ve anlaşılır komut satırı arayüzü

---

## 📋 İçindekiler

- [Kurulum](#-kurulum)
- [Hızlı Başlangıç](#-hızlı-başlangıç)
- [Kullanım](#-kullanım)
- [API Key Alma](#-api-key-alma)
- [Örnekler](#-örnekler)
- [Proje Yapısı](#-proje-yapısı)
- [Geliştirme](#-geliştirme)
- [Katkıda Bulunma](#-katkıda-bulunma)

---

## 🚀 Kurulum

### Gereksinimler

- Go 1.21 veya üzeri
- Financial Modeling Prep API anahtarı ([Nasıl alınır?](#-api-key-alma))

### Adım 1: Projeyi İndirin

```bash
git clone https://github.com/mr-isik/fmp-scraper.git
cd fmp-scraper
```

### Adım 2: Bağımlılıkları Yükleyin

```bash
go mod download
```

### Adım 3: API Anahtarını Ayarlayın

`.env.example` dosyasını `.env` olarak kopyalayın:

```bash
cp .env.example .env
```

`.env` dosyasını düzenleyin ve API anahtarınızı ekleyin:

```env
FMP_API_KEY=your_actual_api_key_here
```

### Adım 4: Uygulamayı Derleyin

```bash
go build -o fmp-scraper.exe ./cmd/fmp-scraper
```

---

## ⚡ Hızlı Başlangıç

Temel kullanım:

```bash
./fmp-scraper -s AAPL -f 2024-01-01 -t 2024-12-31
```

Bu komut Apple (AAPL) hissesinin 2024 yılı boyunca verilerini çeker ve `AAPL_2024-01-01_2024-12-31.csv` dosyasına kaydeder.

---

## 📖 Kullanım

### Temel Komut Yapısı

```bash
fmp-scraper [flags]
```

### Flags (Parametreler)

| Flag       | Kısa | Açıklama                      | Zorunlu  | Örnek                  |
| ---------- | ---- | ----------------------------- | -------- | ---------------------- |
| `--symbol` | `-s` | Hisse senedi sembolü          | ✅ Evet  | `AAPL`, `TSLA`, `MSFT` |
| `--from`   | `-f` | Başlangıç tarihi (YYYY-MM-DD) | ✅ Evet  | `2024-01-01`           |
| `--to`     | `-t` | Bitiş tarihi (YYYY-MM-DD)     | ✅ Evet  | `2024-12-31`           |
| `--output` | `-o` | Çıktı dosya adı               | ❌ Hayır | `my_data.csv`          |

### Çıktı Dosyası

Eğer `--output` parametresi belirtilmezse, dosya adı otomatik olarak şu formatta oluşturulur:

```
{SYMBOL}_{FROM_DATE}_{TO_DATE}.csv
```

Örnek: `AAPL_2024-01-01_2024-12-31.csv`

---

## 🔑 API Key Alma

Financial Modeling Prep API anahtarı almak için:

1. [Financial Modeling Prep](https://site.financialmodelingprep.com/developer/docs) sitesine gidin
2. Ücretsiz hesap oluşturun
3. Dashboard'dan API anahtarınızı kopyalayın
4. `.env` dosyasına yapıştırın

**Ücretsiz Plan Limitleri:**

- 250 API çağrısı/gün
- Temel finansal veriler
- Geçmiş verilere erişim

**Premium planlar** daha yüksek limitler ve ek özellikler sunar.

---

## 💡 Örnekler

### Örnek 1: Tesla Verileri (Varsayılan Dosya Adı)

```bash
fmp-scraper -s TSLA -f 2024-01-01 -t 2024-03-31
```

**Çıktı:** `TSLA_2024-01-01_2024-03-31.csv`

### Örnek 2: Microsoft Verileri (Özel Dosya Adı)

```bash
fmp-scraper -s MSFT -f 2023-06-01 -t 2023-12-31 -o microsoft_2023_h2.csv
```

**Çıktı:** `microsoft_2023_h2.csv`

### Örnek 3: Son 30 Gün

```bash
fmp-scraper -s GOOGL -f 2024-11-01 -t 2024-11-30 -o google_november.csv
```

**Çıktı:** `google_november.csv`

### Örnek 4: Birden Fazla Hisse (Script ile)

PowerShell kullanarak:

```powershell
$symbols = @("AAPL", "TSLA", "MSFT", "GOOGL")
foreach ($symbol in $symbols) {
    ./fmp-scraper -s $symbol -f 2024-01-01 -t 2024-12-31
}
```

---

## 📊 CSV Çıktı Formatı

Oluşturulan CSV dosyası aşağıdaki sütunları içerir:

| Sütun               | Açıklama                       |
| ------------------- | ------------------------------ |
| `Date`              | İşlem tarihi                   |
| `Open`              | Açılış fiyatı                  |
| `High`              | En yüksek fiyat                |
| `Low`               | En düşük fiyat                 |
| `Close`             | Kapanış fiyatı                 |
| `Adjusted Close`    | Düzeltilmiş kapanış fiyatı     |
| `Volume`            | İşlem hacmi                    |
| `Unadjusted Volume` | Düzeltilmemiş hacim            |
| `Change`            | Fiyat değişimi                 |
| `Change Percent`    | Yüzdesel değişim               |
| `VWAP`              | Hacim ağırlıklı ortalama fiyat |
| `Label`             | Etiket bilgisi                 |
| `Change Over Time`  | Zaman içindeki değişim         |

**Örnek Satır:**

```csv
Date,Open,High,Low,Close,Adjusted Close,Volume,...
2024-01-02,184.35,186.95,183.89,185.64,185.64,52164400,...
```

---

## 🏗️ Proje Yapısı

```
fmp-scraper/
├── cmd/
│   └── fmp-scraper/
│       └── main.go              # Uygulama giriş noktası
├── internal/
│   ├── api/
│   │   └── client.go            # FMP API istemcisi
│   ├── cmd/
│   │   └── root.go              # CLI komutları (Cobra)
│   ├── config/
│   │   └── config.go            # Konfigürasyon yönetimi
│   ├── exporter/
│   │   └── csv.go               # CSV export işlemleri
│   └── models/
│       └── stock.go             # Veri modelleri
├── pkg/
│   └── logger/
│       └── logger.go            # Logging utility
├── .env.example                 # Örnek environment dosyası
├── .gitignore
├── go.mod                       # Go modül dosyası
├── go.sum                       # Bağımlılık checksums
└── README.md                    # Bu dosya
```

### Mimari Kararlar (SOLID Prensipler)

#### 1. Single Responsibility Principle (SRP)

- Her paket tek bir sorumluluğa sahip
- `api`: Sadece API iletişimi
- `exporter`: Sadece CSV export
- `config`: Sadece konfigürasyon

#### 2. Open/Closed Principle (OCP)

- Interface'ler kullanılarak genişletilebilir yapı
- `Client`, `Exporter`, `Logger` interface'leri

#### 3. Liskov Substitution Principle (LSP)

- Interface implementasyonları değiştirilebilir

#### 4. Interface Segregation Principle (ISP)

- Küçük, odaklanmış interface'ler
- `HTTPClient`, `Logger`, `Manager`

#### 5. Dependency Inversion Principle (DIP)

- Yüksek seviye modüller düşük seviyeye bağımlı değil
- Dependency injection kullanımı

---

## 🛠️ Geliştirme

### Projeyi Çalıştırma (Development)

```bash
go run ./cmd/fmp-scraper -s AAPL -f 2024-01-01 -t 2024-01-31
```

### Test Ekleme (Gelecek Geliştirmeler)

```bash
go test ./...
```

### Linting

```bash
golangci-lint run
```

### Build Optimizasyonu

Küçük boyutlu binary oluşturmak için:

```bash
go build -ldflags="-s -w" -o fmp-scraper.exe ./cmd/fmp-scraper
```

---

## 🎯 Gelecek Geliştirmeler

- [ ] Birim testleri ekleme
- [ ] JSON ve Excel export desteği
- [ ] Birden fazla sembol için toplu işlem
- [ ] Progress bar ekleme
- [ ] Rate limiting desteği
- [ ] Veri validasyonu ve temizleme
- [ ] Interactive mode
- [ ] Web UI (opsiyonel)

---

## 🤝 Katkıda Bulunma

Katkılarınızı memnuniyetle karşılarız!

1. Fork yapın
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Değişikliklerinizi commit edin (`git commit -m 'Add amazing feature'`)
4. Branch'inizi push edin (`git push origin feature/amazing-feature`)
5. Pull Request açın

---

## 📄 Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Detaylar için [LICENSE](LICENSE) dosyasına bakın.

---

## 👨‍💻 Geliştirici

**Ömer Faruk Işık**

- GitHub: [@mr-isik](https://github.com/mr-isik)

---

## 🙏 Teşekkürler

- [Financial Modeling Prep](https://financialmodelingprep.com/) - API sağlayıcısı
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [godotenv](https://github.com/joho/godotenv) - Environment yönetimi

---

## ⚠️ Sorumluluk Reddi

Bu araç sadece eğitim ve araştırma amaçlıdır. Finansal kararlar almadan önce profesyonel danışmanlık alın. Veri doğruluğu garanti edilmez.

---

<div align="center">

⭐ Bu projeyi beğendiyseniz yıldız vermeyi unutmayın!

**[⬆ Başa Dön](#fmp-scraper)**

</div>
