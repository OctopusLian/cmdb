/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-23 19:09:29
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-23 19:09:30
 */
package utils

import (
	"fmt"
	"time"
)

func TencentAPISignature(secretId, secretKey, host, service, body string, now time.Time) string {

	var (
		requestMethod = "POST"
		uri           = "/"
		queryString   = ""
		headers       = fmt.Sprintf("content-type:application/json; charset=utf-8\nhost:%s\n", host)
		signedHeaders = "content-type;host"
	)

	// 计算格式化request的sha256 hash值
	hashedRequest := Sha256Hex(fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", requestMethod, uri, queryString, headers, signedHeaders, Sha256Hex(body)))

	var (
		algorithm        = "TC3-HMAC-SHA256"
		requestTimestamp = now.Unix()
		credentialScope  = fmt.Sprintf("%s/%s/tc3_request", now.UTC().Format("2006-01-02"), service)
	)

	// 计算签名值
	stringToSign := fmt.Sprintf("%s\n%d\n%s\n%s", algorithm, requestTimestamp, credentialScope, hashedRequest)

	var (
		secretDate    = HS256(now.UTC().Format("2006-01-02"), fmt.Sprintf("TC3%s", secretKey))
		secretService = HS256(service, secretDate)
		secretSinging = HS256("tc3_request", secretService)
		signature     = HS256Hex(stringToSign, secretSinging)
	)

	//计算认证信息
	return fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s", algorithm, secretId, credentialScope, signedHeaders, signature)

}
