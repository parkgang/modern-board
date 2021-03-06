<p align="center">
  <table align="center">
    <tbody>
      <tr>
        <td>
          <img src="https://user-images.githubusercontent.com/63892989/132948190-0e2e9cb6-bafc-4e9f-9bce-c0248397f095.png" alt="screenshot" width="500" />
        </td>
        <td>
          <img src="https://user-images.githubusercontent.com/63892989/132948191-154b0c14-0648-4817-9ca6-32702d03c75f.png" alt="screenshot" width="500" />
        </td>
      </tr>
    </tbody>
  </table>
  <h1 align="center">
    โจ modern-board ๐ค
  </h1>
</p>
<div align="center">

`Bulletins created using modern stacks`

์ผ๋ฐ์ ์ธ ๊ธฐ๋ฅ์ ๊ฐ์ง๊ณ  ์๋ ๊ฒ์ํ **์น ์ฑ**์ด์ง๋ง ๋๋ฆ๋๋ก ํ๋์ ์ธ ๊ธฐ์ ์ ์ฌ์ฉํ์ฌ ์ฌํด์ํ ํ๋ก์ ํธ์๋๋ค. ๋ฃฉ์คํ์ [Microsoft Teams](https://www.microsoft.com/ko-kr/microsoft-teams/group-chat-software) ๋ฅผ ๋ฐ๋ผ๊ฐ๋๋ค. ํ๋ก์ ํธ์ ์ฌ์ฉ ๋ ๊ธฐ์ ์ ๊ฐ ์ ํ์ [webapp (react)](./webapp), [server (gin)](./server) ๋ฅผ ์ฐธ๊ณ ํด์ฃผ์ธ์.

</div>

## Common tool

1. vscode
1. datagrip
1. docker

## Quick start

`docker-compose up -d` ์ผ๋ก docker ์ปจํ์ด๋๋ฅผ ์ฌ๋ฆฌ๊ณ  vscode `์คํ ๋ฐ ๋๋ฒ๊ทธ` ์์ `Server/Client` ์ผ๋ก ์คํํ์ธ์!

๋ฌผ๋ก , ํ์ํ ์์กด์ฑ๊ณผ docker, vscode extension์ ๋ฏธ๋ฆฌ ์ค์น๋์ด ์์ด์ผ ํฉ๋๋ค! ๊ฐ ์ ํ์ `README.md` ๋ฅผ ์ฐธ๊ณ ํ์ธ์.

์ปจํ์ด๋ ์คํ๊ณผ ๊ฐ์ ๋ช๋ น์ด๋ vscode `tasks.json` ์๋ ์ ์๋์ด ์์์ผ๋ก ํธํ๊ฒ task๋ก ์คํํ์ธ์! ๐

> ์ฒซ ์ปจํ์ด๋๋ฅผ ํ๋ก๋น์ ๋ํ๋๋ฐ ์๊ฐ์ด ์์๋๋ฏ๋ก ์๋์ ๋ช๋ น์ด๋ก 1ํ ์ปจํ์ด๋ apply ์ดํ ์ฌ์ ๋ฅผ ์๊ฐ์ ๋๊ณ  ๋๋ฒ๊น์ ์คํ์์ผ์ฃผ์ธ์.

## Prod deployment

prod์ผ๋ก ์ฝ๊ฒ ๋ฐฐํฌํ  ์ ์๋๋ก `Dockerfile` ํ์ผ์ด ๋ชจ๋ ์์ฑ๋์ด ์์ผ๋ ๊ฐ๋ฐ ํ๋ก์ธ์ค์ ๋ง์ถฐ์ `CI/CD` pipeline์ผ๋ก ํ์ฐ๊ฑฐ๋ ํน์ ์๋์ cli๋ก ๋ฐฐํฌ์ ์ฌ๋ฏธ๋ฅผ ๋๊ปด๋ณด์ธ์!

```shell
docker build -t parkgang/modern-board:0.1.0 .
docker run --name modern-board-prod -p 22000:8080 -d parkgang/modern-board:0.1.0
```
