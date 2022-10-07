# ExifRemoverBot
- A telegram bot that can remove exif from JPEG and PNG files
- You send photo to bot and bot send you that photo without additional EXIF metadata

## Example
- Original EXIF data:
>- File Name                       : Alan_Turing_Aged_16.jpg
>- Directory                       : .
>- File Size                       : 217 kB
>- File Modification Date/Time     : 2022:03:10 10:51:43+03:30
>- File Access Date/Time           : 2022:10:08 01:26:55+03:30
>- File Permissions                : -rwxrwxrwx
>- File Inode Change Date/Time     : 2022:10:08 01:26:48+03:30
>- File Type                       : JPEG
>- File Type Extension             : jpg
>- MIME Type                       : image/jpeg
>- JFIF Version                    : 1.01
>- Exif Byte Order                 : Big-endian (Motorola, MM)
>- X Resolution                    : 72
>- Y Resolution                    : 72
>- Resolution Unit                 : inches
>- Software                        : GIMP 2.8.16
>- Modify Date                     : 2017:09:02 20:36:29
>- Artist                          : Stas
>- Exif Version                    : 0220
>- Date/Time Original              : 2012:11:23 18:51:50
>- Create Date                     : 2012:11:23 18:51:50
>- Sub Sec Time Original           : 14
>- Sub Sec Time Digitized          : 14
>- Flashpix Version                : 0100
>- Color Space                     : Uncalibrated
>- Exif Image Width                : 675
>- Exif Image Height               : 919
>- Interoperability Version        : 0100
>- Related Image Width             : 752
>- Related Image Height            : 1043
>- Image Unique ID                 : dec1c8fc87342797a808e3273b191439
>- XP Author                       : Stas
>- Padding                         : (Binary data 2060 bytes, use -b option to extract)
>- Compression                     : JPEG (old-style)
>- Thumbnail Offset                : 4696
>- Thumbnail Length                : 4145
>- Creator                         : Stas
>- Instance ID                     : uuid:faf5bdd5-ba3d-11da-ad31-d33d75182f1b
>- Date Time                       : 2015:07:26 02:49:29
>- Flash Pix Version               : FlashPix Version 1.0
>- Interoperability Version        : 0100
>- Related Image Length            : 1043
>- Date/Time Digitized             : 2012:11:23 18:51:50
>- Image Width                     : 675
>- Image Height                    : 919
>- Encoding Process                : Baseline DCT, Huffman coding
>- Bits Per Sample                 : 8
>- Color Components                : 1
>- Image Size                      : 675x919
>- Megapixels                      : 0.620
>- Create Date                     : 2012:11:23 18:51:50.14
>- Date/Time Original              : 2012:11:23 18:51:50.14
>- Thumbnail Image                 : (Binary data 4145 bytes, use -b option to extract)

- After Remove EXIF data:
>- File Name                       : file_32.jpg
>- Directory                       : .
>- File Size                       : 217 kB
>- File Modification Date/Time     : 2022:10:08 01:30:18+03:30
>- File Access Date/Time           : 2022:10:08 01:30:35+03:30
>- File Inode Change Date/Time     : 2022:10:08 01:30:18+03:30
>- File Permissions                : -rw-r--r--
>- File Type                       : JPEG
>- File Type Extension             : jpg
>- MIME Type                       : image/jpeg
>- JFIF Version                    : 1.01
>- Warning                         : Malformed APP1 EXIF segment
>- Create Date                     : 2012:11:23 18:51:50.143
>- Modify Date                     : 2015:07:26 02:49:29+05:00
>- Creator                         : Stas
>- Instance ID                     : uuid:faf5bdd5-ba3d-11da-ad31-d33d75182f1b
>- X Resolution                    : 72
>- Y Resolution                    : 72
>- Resolution Unit                 : Inch
>- Software                        : Picasa
>- Date Time                       : 2015:07:26 02:49:29
>- Artist                          : Stas
>- XP Author                       : Stas
>- Padding                         : 2060 bytes undefined data
>- Compression                     : JPEG compression
>- Sub Sec Time Original           : 14
>- Sub Sec Time Digitized          : 14
>- Flash Pix Version               : FlashPix Version 1.0
>- Interoperability Version        : 0100
>- Related Image Width             : 752
>- Related Image Length            : 1043
>- Exif Version                    : Exif Version 2.2
>- Date/Time Original              : 2012:11:23 18:51:50
>- Date/Time Digitized             : 2012:11:23 18:51:50
>- Color Space                     : Unknown (Internal error (unknown value 65535))
>- Exif Image Width                : 707
>- Exif Image Height               : 919
>- Image Unique ID                 : dec1c8fc87342797a808e3273b191439
>- Image Width                     : 675
>- Image Height                    : 919
>- Encoding Process                : Baseline DCT, Huffman coding
>- Bits Per Sample                 : 8
>- Color Components                : 1
>- Image Size                      : 675x919
>- Megapixels                      : 0.620
