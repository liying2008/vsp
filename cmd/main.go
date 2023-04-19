package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "vsp",
	})

	app.All("/*", func(c *fiber.Ctx) error {
		targetUrl := c.Params("*")
		log.Println("targetUrl", targetUrl)
		if targetUrl == "favicon.ico" {
			// TODO
			return nil
		}

		client := &fasthttp.Client{}

		// 获取 GoFiber 的请求和响应对象
		ctx := c.Context()

		// 创建 fasthttp.Request 对象，并将 GoFiber 的请求参数复制到该对象中
		req := &fasthttp.Request{}
		ctx.Request.Header.CopyTo(&req.Header)
		req.Header.SetMethod(string(ctx.Method()))
		req.Header.SetRequestURI(targetUrl)

		// 发送 fasthttp.Request 对象，并获取 fasthttp.Response 对象
		resp := &fasthttp.Response{}
		err := client.Do(req, resp)
		if err != nil {
			return err
		}

		// 将 fasthttp.Response 对象的内容复制到 GoFiber 的响应对象中
		resp.Header.CopyTo(&ctx.Response.Header)
		ctx.Response.SetStatusCode(resp.StatusCode())
		ctx.Response.SetBody(resp.Body())

		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
