Creating the migrations: 

```
 goose -dir "db/migrations" create create_user_table sql

```

Apply the migrations:

```

goose -dir "db/migrations" mysql "root:Jyoti@123@tcp(127.0.0.1:3306)/auth_dev" up

```


### Gmake installation steps:


Got it! Let’s simplify the README to **just installing and using GNU Make on Windows via MSYS2**, without the VS Code-specific stuff.

Here’s the cleaned-up version:

---

# Installing and Using GNU Make on Windows (via MSYS2)

This guide explains how to install **GNU Make (`gmake`)** on Windows using **MSYS2** and make it available globally.

---

## 1. Install MSYS2

1. Download MSYS2 from [https://www.msys2.org/](https://www.msys2.org/).
2. Install it (default path: `C:\msys64`).
3. Open the **MSYS2 MinGW64** terminal from the Start Menu.

---

## 2. Update MSYS2

Run the following in the MSYS2 MinGW64 terminal:

```sh
pacman -Syu
# Close and reopen the terminal, then run:
pacman -Su
```

---

## 3. Install GNU Make

```sh
pacman -S make
```

Verify installation:

```sh
make --version
```

Expected output:

```
GNU Make 4.x
```

---

## 4. Add Make to Windows PATH

1. Open **System Properties**:
   Press `Win + R`, type `sysdm.cpl`, hit Enter.
2. Go to **Advanced → Environment Variables**.
3. Under **System variables**, find **Path**, click **Edit → New**.
4. Add:

```
C:\msys64\usr\bin
```

5. Save and close all dialogs.
6. Open a new PowerShell or CMD window.

Test:

```powershell
make --version
```

✅ If you see the version, `make` is now available globally on your system.

---

## 5. Troubleshooting

* If `make --version` still fails:

  * Run the full path directly:

    ```powershell
    C:\msys64\usr\bin\make.exe --version
    ```
  * Ensure `C:\msys64\usr\bin` was correctly added to PATH.
* If `make.exe` is not in `/usr/bin`, it may be in `/mingw64/bin`. Check in MSYS2 terminal:

  ```sh
  where make
  ```

---
