package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"net/url"
)

func main() {
	app := fiber.New()

	app.Get("/:url", func(c *fiber.Ctx) error {
		targetUrl := c.Params("url")
		println("targetUrl", targetUrl)
		target, err := url.Parse(targetUrl)
		if err != nil {
			return err
		}

		println("target", target.Host, target.RequestURI())
		// 创建 fasthttp.HostClient 对象来处理 HTTP 请求
		client := &fasthttp.HostClient{
			Addr: target.Host,
		}

		// 获取 GoFiber 的请求和响应对象
		ctx := c.Context()

		// 创建 fasthttp.Request 对象，并将 GoFiber 的请求参数复制到该对象中
		req := &fasthttp.Request{}
		ctx.Request.Header.CopyTo(&req.Header)
		req.Header.SetMethod(string(ctx.Method()))
		req.Header.SetRequestURI(targetUrl)

		// 发送 fasthttp.Request 对象，并获取 fasthttp.Response 对象
		resp := &fasthttp.Response{}
		err = client.Do(req, resp)
		if err != nil {
			return err
		}

		// 将 fasthttp.Response 对象的内容复制到 GoFiber 的响应对象中
		resp.Header.CopyTo(&ctx.Response.Header)
		ctx.Response.SetStatusCode(resp.StatusCode())
		ctx.Response.SetBody(resp.Body())

		return nil
	})

	app.Listen(":3000")
}
