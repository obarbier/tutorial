package main

import (
	"encoding/pem"
	"log"
	"testing"
)

var edTest = map[string]string{
	"private": "-----BEGIN ED25519 PRIVATE KEY-----\n1sj6meD/iKe4QGnJrJHoCqlz+3JOaRTfD+/V7G1A6By/fVo5m1FF/XVSTkScjoF8\n9EfICjIQfdy/3NrLblVveg==\n-----END ED25519 PRIVATE KEY-----\n",
	"public":  "-----BEGIN ED25519 PUBLIC KEY-----\nv31aOZtRRf11Uk5EnI6BfPRHyAoyEH3cv9zay25Vb3o=\n-----END ED25519 PUBLIC KEY-----",
}

var rsaTest = map[string]string{
	"private": "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAnJ40wjF1Q1vI2tNTZ5OXOobDy4bjW8qpja2SBjaDa9NTIzd7\ndLfLQRmjP1QFO+6pir+NUwdQUKXMUIsfYl1QgcNVJm2qz3ex3DbIKGbi9CMttu9C\nJOrg/ao6buamwGJ0ePRjEO9WWuBRu/s0p3bI2c7zUaheP+DUSVqrGixXpTJjt+9z\nD97rxD6qkl7ZEZorXrQRmKnf7Hi7ZPat4gjTOAB0BgYwbXD9yxrY1dr2x4b9WoDs\nhgj7DUoJr+cb7V2DepyastwUia6DC6NsBjEuBbO0BGC+JKbn+355zXwF2QUExx7Z\nuFWZ4cN3UZWsF6VETrFqpWWz8OEiwetlQVcTPQIDAQABAoIBAF/v5V+DTlJ5mdq5\nkqCi3wNB5BP7R7BFv4EC7q0RnYViSM4MwXooz7/MBZzYSfCBbKeKWPagR0lvlm1M\nG2h0wskKL1G/4d7+chv0Dr348FMebXVesETPAA1CxlKCxWiZpsEk6r5H7bzzJf4h\njgp7D+OkCpZdrYYxobhhaug4e2O8OdWigXV6Cor4drmomF7Am+4rxEkcW+lqZ0wp\noPfE8zuzvSpSRu5Gi3i44pTv7krrD+9AGhBrETFbpVCN1Gp3GNTIBFGXwmfB5VAB\nvGCu0Jvds5Ky0oYtL6CFHThgYW9VEcTxwdtQ3Ga7nyKg8qvlhR0PNUULjxRCKlPJ\nVDdoyV0CgYEAxamgTmO1hImdKLKEWDkE7Xn3qpQOElphr66psvH2fdsIXlGmh8Tv\nvhX3MIJ8HQIlMwhWlzBpV0KloDEe+TGpGnz0fm4ynaRZ/1fVeERYr0OCM3+GX36T\niSFIuvcJon/JtNWJKKGv8iTl57CN2RUED/JLJAO1MYq2V44c8ZVN5dsCgYEAytd1\nm6cLVgsScaG4tKSrPNe8Q/2VY0G2uYk10YycpEQ9XTVMdnmZkMZzQ7zisplKOooT\n4Nnak9PYWPcPx7hvvfAmwtW8TuJK9iE5ott+BkGyemNidRbIVcEjP25IcTAv1U9L\nypOPLbMVS703Qo5DWxpqdQCv4JSYf3iGpEjgEscCgYAlaHPZQg+RVMX0dMyNMcVX\n+DRCCSEcohRIvmKJZjeDHBfaWdONcFz6+Yc9nARHLSfDH7nbhSL6i7dyuLkm6hoZ\n1DolT0+u+/K4W3Qf/bdW/AzBGEpi+j6LvkvYbnZZVZvj6GG72dXFmuwTzBscUVji\nd7V2zGjXRmw558BcjCgg7wKBgQC98cll+qT46+FXmzNVpaXckcDwgjQ0AENVtpE5\nK0073dMAx3pUr0YIdm1VjzlSOY1zB/3ZXuf56jT8Ck8ynRm2aX5CkNeHnwzHMxG/\n2X5H6ZlyVkLKloz/EJ8rNOCNbaw9OcYlasx25O4aQJPQtVuJolDyQ7Ruv3eWVozu\npnd93QKBgGQ24OoLFOyeWTuGiDlsfGzZJzwsm54YTtHRqxdYfosAEWUK8udIdPXl\ntcIxX7/ta1C3nfMLApDzJie1Vg48+XFmlOMF3JDDVoRmgvFM7Vz/swsYg66dULxv\nvucFkSN8+rsYHfjGcXJ2r9KLwU2V/sBABIMOkqpQivFgP+9xhkrk\n-----END RSA PRIVATE KEY-----\n",
	"public":  "-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEAnJ40wjF1Q1vI2tNTZ5OXOobDy4bjW8qpja2SBjaDa9NTIzd7dLfL\nQRmjP1QFO+6pir+NUwdQUKXMUIsfYl1QgcNVJm2qz3ex3DbIKGbi9CMttu9CJOrg\n/ao6buamwGJ0ePRjEO9WWuBRu/s0p3bI2c7zUaheP+DUSVqrGixXpTJjt+9zD97r\nxD6qkl7ZEZorXrQRmKnf7Hi7ZPat4gjTOAB0BgYwbXD9yxrY1dr2x4b9WoDshgj7\nDUoJr+cb7V2DepyastwUia6DC6NsBjEuBbO0BGC+JKbn+355zXwF2QUExx7ZuFWZ\n4cN3UZWsF6VETrFqpWWz8OEiwetlQVcTPQIDAQAB\n-----END RSA PUBLIC KEY-----",
}

func getpriv() []byte {
	b := edTest["private"]
	p, _ := pem.Decode([]byte(b))
	return p.Bytes
}

func getpub() []byte {
	b := edTest["public"]
	p, _ := pem.Decode([]byte(b))
	return p.Bytes
}

func getRsapriv() []byte {
	b := rsaTest["private"]
	p, _ := pem.Decode([]byte(b))
	return p.Bytes
}

func getRsapub() []byte {
	b := rsaTest["public"]
	p, _ := pem.Decode([]byte(b))
	return p.Bytes
}

func Benchmark_GenerateEd25519(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			generate(ED25519, mockWriter{}, mockWriter{})
		}
	})
}

func Benchmark_Sign_and_verify_ED(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			msg := []byte("Hello World")
			sig := Sign(ED25519, getpriv(), msg)
			verify := Verify(ED25519, getpub(), msg, []byte(sig))
			if !verify {
				log.Fatal("failed to verify")
			}
		}
	})
}

func Benchmark_Sign_and_verify_RSA(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			msg := []byte("Hello World")
			sig := Sign(RSA, getRsapriv(), msg)
			verify := Verify(RSA, getRsapub(), msg, []byte(sig))
			if !verify {
				log.Fatal("failed to verify")
			}
		}
	})
}

func Benchmark_GenerateRsa(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			generate(RSA, mockWriter{}, mockWriter{})
		}
	})
}

type mockWriter struct {
}

func (r mockWriter) Write(p []byte) (n int, err error) {
	return
}
