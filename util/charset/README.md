# character 说明

character描述了如下内容:
- mysql的字符集
- mysql的文字排序方法
- mysql的编码转换方法

## 字符集

mysql包含40种字符集

字符集|
---|
 armscii8           |
 ascii              |
 big5               |
 binary             |
 cp1250             |
 cp1251             |
 cp1256             |
 cp1257             |
 cp850              |
 cp852              |
 cp866              |
 cp932              |
 dec8               |
 eucjpms            |
 euckr              |
 gb2312             |
 gbk                |
 geostd8            |
 greek              |
 hebrew             |
 hp8                |
 keybcs2            |
 koi8r              |
 koi8u              |
 latin1             |
 latin2             |
 latin5             |
 latin7             |
 macce              |
 macroman           |
 sjis               |
 swe7               |
 tis620             |
 ucs2               |
 ujis               |
 utf16              |
 utf16le            |
 utf32              |
 utf8               |
 utf8mb4            |


## 文字排序方法

Collation | CharSet
---|---
 armscii8_bin             | armscii8     
 armscii8_general_ci      | armscii8     
 ascii_bin                | ascii        
 ascii_general_ci         | ascii        
 big5_bin                 | big5         
 big5_chinese_ci          | big5         
 binary                   | binary       
 cp1250_croatian_ci       | cp1250       
 cp1250_czech_cs          | cp1250       
 cp1250_general_ci        | cp1250       
 cp1250_polish_ci         | cp1250       
 cp1250_bin               | cp1250       
 cp1251_general_cs        | cp1251       
 cp1251_general_ci        | cp1251       
 cp1251_bin               | cp1251       
 cp1251_ukrainian_ci      | cp1251       
 cp1251_bulgarian_ci      | cp1251       
 cp1256_bin               | cp1256       
 cp1256_general_ci        | cp1256       
 cp1257_bin               | cp1257       
 cp1257_lithuanian_ci     | cp1257       
 cp1257_general_ci        | cp1257       
 cp850_bin                | cp850        
 cp850_general_ci         | cp850        
 cp852_bin                | cp852        
 cp852_general_ci         | cp852        
 cp866_bin                | cp866        
 cp866_general_ci         | cp866        
 cp932_bin                | cp932        
 cp932_japanese_ci        | cp932        
 dec8_swedish_ci          | dec8         
 dec8_bin                 | dec8         
 eucjpms_bin              | eucjpms      
 eucjpms_japanese_ci      | eucjpms      
 euckr_bin                | euckr        
 euckr_korean_ci          | euckr        
 gb2312_bin               | gb2312       
 gb2312_chinese_ci        | gb2312       
 gbk_bin                  | gbk          
 gbk_chinese_ci           | gbk          
 geostd8_bin              | geostd8      
 geostd8_general_ci       | geostd8      
 greek_bin                | greek        
 greek_general_ci         | greek        
 hebrew_bin               | hebrew       
 hebrew_general_ci        | hebrew       
 hp8_bin                  | hp8          
 hp8_english_ci           | hp8          
 keybcs2_bin              | keybcs2      
 keybcs2_general_ci       | keybcs2      
 koi8r_bin                | koi8r        
 koi8r_general_ci         | koi8r        
 koi8u_bin                | koi8u        
 koi8u_general_ci         | koi8u        
 latin1_german2_ci        | latin1       
 latin1_danish_ci         | latin1       
 latin1_swedish_ci        | latin1       
 latin1_german1_ci        | latin1       
 latin1_spanish_ci        | latin1       
 latin1_general_cs        | latin1       
 latin1_general_ci        | latin1       
 latin1_bin               | latin1       
 latin2_bin               | latin2       
 latin2_croatian_ci       | latin2       
 latin2_hungarian_ci      | latin2       
 latin2_general_ci        | latin2       
 latin2_czech_cs          | latin2       
 latin5_bin               | latin5       
 latin5_turkish_ci        | latin5       
 latin7_general_cs        | latin7       
 latin7_general_ci        | latin7       
 latin7_estonian_cs       | latin7       
 latin7_bin               | latin7       
 macce_bin                | macce        
 macce_general_ci         | macce        
 macroman_bin             | macroman     
 macroman_general_ci      | macroman     
 sjis_bin                 | sjis         
 sjis_japanese_ci         | sjis         
 swe7_bin                 | swe7         
 swe7_swedish_ci          | swe7         
 tis620_bin               | tis620       
 tis620_thai_ci           | tis620       
 ucs2_polish_ci           | ucs2         
 ucs2_roman_ci            | ucs2         
 ucs2_slovenian_ci        | ucs2         
 ucs2_spanish2_ci         | ucs2         
 ucs2_romanian_ci         | ucs2         
 ucs2_general_mysql500_ci | ucs2         
 ucs2_slovak_ci           | ucs2         
 ucs2_latvian_ci          | ucs2         
 ucs2_vietnamese_ci       | ucs2         
 ucs2_lithuanian_ci       | ucs2         
 ucs2_icelandic_ci        | ucs2         
 ucs2_unicode_520_ci      | ucs2         
 ucs2_danish_ci           | ucs2         
 ucs2_unicode_ci          | ucs2         
 ucs2_croatian_ci         | ucs2         
 ucs2_czech_ci            | ucs2         
 ucs2_bin                 | ucs2         
 ucs2_german2_ci          | ucs2         
 ucs2_turkish_ci          | ucs2         
 ucs2_general_ci          | ucs2         
 ucs2_sinhala_ci          | ucs2         
 ucs2_swedish_ci          | ucs2         
 ucs2_hungarian_ci        | ucs2         
 ucs2_spanish_ci          | ucs2         
 ucs2_esperanto_ci        | ucs2         
 ucs2_estonian_ci         | ucs2         
 ucs2_persian_ci          | ucs2         
 ujis_bin                 | ujis         
 ujis_japanese_ci         | ujis         
 utf16_hungarian_ci       | utf16        
 utf16_spanish_ci         | utf16        
 utf16_esperanto_ci       | utf16        
 utf16_estonian_ci        | utf16        
 utf16_persian_ci         | utf16        
 utf16_polish_ci          | utf16        
 utf16_roman_ci           | utf16        
 utf16_slovenian_ci       | utf16        
 utf16_spanish2_ci        | utf16        
 utf16_romanian_ci        | utf16        
 utf16_slovak_ci          | utf16        
 utf16_latvian_ci         | utf16        
 utf16_vietnamese_ci      | utf16        
 utf16_lithuanian_ci      | utf16        
 utf16_icelandic_ci       | utf16        
 utf16_unicode_520_ci     | utf16        
 utf16_danish_ci          | utf16        
 utf16_unicode_ci         | utf16        
 utf16_croatian_ci        | utf16        
 utf16_czech_ci           | utf16        
 utf16_bin                | utf16        
 utf16_german2_ci         | utf16        
 utf16_turkish_ci         | utf16        
 utf16_general_ci         | utf16        
 utf16_sinhala_ci         | utf16        
 utf16_swedish_ci         | utf16        
 utf16le_bin              | utf16le      
 utf16le_general_ci       | utf16le      
 utf32_esperanto_ci       | utf32        
 utf32_estonian_ci        | utf32        
 utf32_persian_ci         | utf32        
 utf32_polish_ci          | utf32        
 utf32_roman_ci           | utf32        
 utf32_slovenian_ci       | utf32        
 utf32_spanish2_ci        | utf32        
 utf32_romanian_ci        | utf32        
 utf32_slovak_ci          | utf32        
 utf32_latvian_ci         | utf32        
 utf32_vietnamese_ci      | utf32        
 utf32_lithuanian_ci      | utf32        
 utf32_icelandic_ci       | utf32        
 utf32_unicode_520_ci     | utf32        
 utf32_danish_ci          | utf32        
 utf32_unicode_ci         | utf32        
 utf32_croatian_ci        | utf32        
 utf32_czech_ci           | utf32        
 utf32_bin                | utf32        
 utf32_german2_ci         | utf32        
 utf32_turkish_ci         | utf32        
 utf32_general_ci         | utf32        
 utf32_sinhala_ci         | utf32        
 utf32_swedish_ci         | utf32        
 utf32_hungarian_ci       | utf32        
 utf32_spanish_ci         | utf32        
 utf8_unicode_ci          | utf8         
 utf8_croatian_ci         | utf8         
 utf8_czech_ci            | utf8         
 utf8_bin                 | utf8         
 utf8_german2_ci          | utf8         
 utf8_turkish_ci          | utf8         
 utf8_general_ci          | utf8         
 utf8_sinhala_ci          | utf8         
 utf8_swedish_ci          | utf8         
 utf8_hungarian_ci        | utf8         
 utf8_spanish_ci          | utf8         
 utf8_esperanto_ci        | utf8         
 utf8_estonian_ci         | utf8         
 utf8_persian_ci          | utf8         
 utf8_polish_ci           | utf8         
 utf8_roman_ci            | utf8         
 utf8_slovenian_ci        | utf8         
 utf8_spanish2_ci         | utf8         
 utf8_romanian_ci         | utf8         
 utf8_general_mysql500_ci | utf8         
 utf8_slovak_ci           | utf8         
 utf8_latvian_ci          | utf8         
 utf8_vietnamese_ci       | utf8         
 utf8_lithuanian_ci       | utf8         
 utf8_icelandic_ci        | utf8         
 utf8_unicode_520_ci      | utf8         
 utf8_danish_ci           | utf8         
 utf8mb4_esperanto_ci     | utf8mb4      
 utf8mb4_estonian_ci      | utf8mb4      
 utf8mb4_persian_ci       | utf8mb4      
 utf8mb4_polish_ci        | utf8mb4      
 utf8mb4_roman_ci         | utf8mb4      
 utf8mb4_slovenian_ci     | utf8mb4      
 utf8mb4_spanish2_ci      | utf8mb4      
 utf8mb4_romanian_ci      | utf8mb4      
 utf8mb4_slovak_ci        | utf8mb4      
 utf8mb4_latvian_ci       | utf8mb4      
 utf8mb4_vietnamese_ci    | utf8mb4      
 utf8mb4_lithuanian_ci    | utf8mb4      
 utf8mb4_icelandic_ci     | utf8mb4      
 utf8mb4_unicode_520_ci   | utf8mb4      
 utf8mb4_danish_ci        | utf8mb4      
 utf8mb4_unicode_ci       | utf8mb4      
 utf8mb4_croatian_ci      | utf8mb4      
 utf8mb4_czech_ci         | utf8mb4      
 utf8mb4_bin              | utf8mb4      
 utf8mb4_german2_ci       | utf8mb4      
 utf8mb4_turkish_ci       | utf8mb4      
 utf8mb4_general_ci       | utf8mb4      
 utf8mb4_sinhala_ci       | utf8mb4      
 utf8mb4_swedish_ci       | utf8mb4      
 utf8mb4_hungarian_ci     | utf8mb4      
 utf8mb4_spanish_ci       | utf8mb4      


## 编码转换方法

提供UTF8和其他编码互相转换的方法