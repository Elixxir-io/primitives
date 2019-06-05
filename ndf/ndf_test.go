package globals

import (
	"testing"
)

var ExampleNDF2 = `{"Timestamp": "2019-06-04 20:48:48.284347", "gateways": [{"Address": "52.25.135.52", "Tls_certificate": "-----BEGIN CERTIFICATE-----MIIDgTCCAmmgAwIBAgIJAKLdZ8UigIAeMA0GCSqGSIb3DQEBBQUAMG8xCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEaMBgGA1UEAwwRZ2F0ZXdheSouY21peC5yaXAwHhcNMTkwMzA1MTgzNTU0WhcNMjkwMzAyMTgzNTU0WjBvMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJQ2xhcmVtb250MRswGQYDVQQKDBJQcml2YXRlZ3JpdHkgQ29ycC4xGjAYBgNVBAMMEWdhdGV3YXkqLmNtaXgucmlwMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9+AaxwDPxHbhLmn4HoZu0oUM48Qufc6T5XEZTrpMrqJAouXk+61Jc0EFH96/sbj7VyvnXPRogIENbk2Y84BkB9SkRMIXya/gh9dOEDSgnvj/yg24l3bdKFqBMKiFg00PYB30fU+Abe3OI/le0I+v++RwH2AV0BMq+T6PcAGjCC1Q1ZB0wP9/VqNMWq5lbK9wD46IQiSi+SgIQeE7HoiAZXrGO0Y7l9P3+VRoXjRQbqfn3ETNL9ZvQuarwAYC9Ix5MxUrS5agOmfjc8bfkpYDFAXRXmdKNISJmtCebX2kDrpP8Bdasx7Fzsx59cEUHCl2aJOWXc7R5m3juOVL1HUxjQIDAQABoyAwHjAcBgNVHREEFTATghFnYXRld2F5Ki5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEAMu3xoc2LW2UExAAIYYWEETggLNrlGonxteSujuJjOR+ik5SVLn0lEu22+z+FCA7gSk9FkWu+v9qnfOfm2Am+WKYWv3dJ5RypW/hDNXkOYxVJNYFxeShnHohNqq4eDKpdqSxEcuErFXJdLbZP1uNs4WIOKnThgzhkpuy7tZRosvOF1X5uL1frVJzHN5jASEDAa7hJNmQ24kh+ds/Ge39fGD8pK31CWhnIXeDovKD7wivi/gSOBtcRWWLvU8SizZkS3hgTw0lSOf5geuzvasCEYlqrKFssj6cTzbCBxy3ra3WazRTNTW4TmkHlCUC9I3oWTTxw5iQxF/I2kQQnwR7L3w==-----END CERTIFICATE-----"}, {"Address": "52.25.219.38", "Tls_certificate": "-----BEGIN CERTIFICATE-----MIIDgTCCAmmgAwIBAgIJAKLdZ8UigIAeMA0GCSqGSIb3DQEBBQUAMG8xCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEaMBgGA1UEAwwRZ2F0ZXdheSouY21peC5yaXAwHhcNMTkwMzA1MTgzNTU0WhcNMjkwMzAyMTgzNTU0WjBvMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJQ2xhcmVtb250MRswGQYDVQQKDBJQcml2YXRlZ3JpdHkgQ29ycC4xGjAYBgNVBAMMEWdhdGV3YXkqLmNtaXgucmlwMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9+AaxwDPxHbhLmn4HoZu0oUM48Qufc6T5XEZTrpMrqJAouXk+61Jc0EFH96/sbj7VyvnXPRogIENbk2Y84BkB9SkRMIXya/gh9dOEDSgnvj/yg24l3bdKFqBMKiFg00PYB30fU+Abe3OI/le0I+v++RwH2AV0BMq+T6PcAGjCC1Q1ZB0wP9/VqNMWq5lbK9wD46IQiSi+SgIQeE7HoiAZXrGO0Y7l9P3+VRoXjRQbqfn3ETNL9ZvQuarwAYC9Ix5MxUrS5agOmfjc8bfkpYDFAXRXmdKNISJmtCebX2kDrpP8Bdasx7Fzsx59cEUHCl2aJOWXc7R5m3juOVL1HUxjQIDAQABoyAwHjAcBgNVHREEFTATghFnYXRld2F5Ki5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEAMu3xoc2LW2UExAAIYYWEETggLNrlGonxteSujuJjOR+ik5SVLn0lEu22+z+FCA7gSk9FkWu+v9qnfOfm2Am+WKYWv3dJ5RypW/hDNXkOYxVJNYFxeShnHohNqq4eDKpdqSxEcuErFXJdLbZP1uNs4WIOKnThgzhkpuy7tZRosvOF1X5uL1frVJzHN5jASEDAa7hJNmQ24kh+ds/Ge39fGD8pK31CWhnIXeDovKD7wivi/gSOBtcRWWLvU8SizZkS3hgTw0lSOf5geuzvasCEYlqrKFssj6cTzbCBxy3ra3WazRTNTW4TmkHlCUC9I3oWTTxw5iQxF/I2kQQnwR7L3w==-----END CERTIFICATE-----"}, {"Address": "52.41.80.104", "Tls_certificate": "-----BEGIN CERTIFICATE-----MIIDgTCCAmmgAwIBAgIJAKLdZ8UigIAeMA0GCSqGSIb3DQEBBQUAMG8xCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEaMBgGA1UEAwwRZ2F0ZXdheSouY21peC5yaXAwHhcNMTkwMzA1MTgzNTU0WhcNMjkwMzAyMTgzNTU0WjBvMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJQ2xhcmVtb250MRswGQYDVQQKDBJQcml2YXRlZ3JpdHkgQ29ycC4xGjAYBgNVBAMMEWdhdGV3YXkqLmNtaXgucmlwMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9+AaxwDPxHbhLmn4HoZu0oUM48Qufc6T5XEZTrpMrqJAouXk+61Jc0EFH96/sbj7VyvnXPRogIENbk2Y84BkB9SkRMIXya/gh9dOEDSgnvj/yg24l3bdKFqBMKiFg00PYB30fU+Abe3OI/le0I+v++RwH2AV0BMq+T6PcAGjCC1Q1ZB0wP9/VqNMWq5lbK9wD46IQiSi+SgIQeE7HoiAZXrGO0Y7l9P3+VRoXjRQbqfn3ETNL9ZvQuarwAYC9Ix5MxUrS5agOmfjc8bfkpYDFAXRXmdKNISJmtCebX2kDrpP8Bdasx7Fzsx59cEUHCl2aJOWXc7R5m3juOVL1HUxjQIDAQABoyAwHjAcBgNVHREEFTATghFnYXRld2F5Ki5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEAMu3xoc2LW2UExAAIYYWEETggLNrlGonxteSujuJjOR+ik5SVLn0lEu22+z+FCA7gSk9FkWu+v9qnfOfm2Am+WKYWv3dJ5RypW/hDNXkOYxVJNYFxeShnHohNqq4eDKpdqSxEcuErFXJdLbZP1uNs4WIOKnThgzhkpuy7tZRosvOF1X5uL1frVJzHN5jASEDAa7hJNmQ24kh+ds/Ge39fGD8pK31CWhnIXeDovKD7wivi/gSOBtcRWWLvU8SizZkS3hgTw0lSOf5geuzvasCEYlqrKFssj6cTzbCBxy3ra3WazRTNTW4TmkHlCUC9I3oWTTxw5iQxF/I2kQQnwR7L3w==-----END CERTIFICATE-----"}], "nodes": [{"Id": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "Dsa_public_key": "-----BEGIN PUBLIC KEY-----\nMIIDNDCCAiwCggEBAJ22+1lRtmu2/h4UDx0s5VAjdBYf1lON8WSCGGQvC1xIyPek\nGq36GHMkuHZ0+hgisA8ez4E2lD18VXVyZOWhpE/+AS6ZNuAMHT6TELAcfReYBdMF\niyqfS7b5cWv+YRfGtbPMTZvjQRBK1KgK1slOAF9LmT4U8JHrUXQ78zBQw43iNVZ+\nGzTD1qXAzqoaDzaCE8PRmEPQtLCdy5/HLTnI3kHxvxTUu0Vjyig3FiHK0zJLai05\nIUW+v6x0iAUjb1yi/pK4cc2PnDbTKStVCcqMqneirfx7/XfdpvcRJadFb+oVPkMy\nVqImHGoG7TaTeX55lfrVqrvPvj7aJ0HjdUBK4lsCIQDywxGTdM52yTVpkLRlN0oX\n8j+e01CJvZafYcbd6ZmMHwKCAQBcf/awb48UP+gohDNJPkdpxNmIrOW+JaDiSAln\nBxbGE9ewzuaTL4+qfETSyyRSPaU/vk9uw1lYktGqWMQyigbEahVmLn6qcDod7Pi7\nstBdvi65VsFCozhmHRBGHA0TVHIIUFfzSUMJ/6c8YR94syrbtXQMNhyfNb6QmX2y\nAU4u9apheC9Sq+uL1kMsTdCXvFQjsoXa+2DcNk6BYfSio1rKOhCxxNIDzHakcKM6\n/cvdkpWYWavYtW4XJSUteOrGbnG6muPx3SSHGZh0OTzU2DIYaABlR2Dh40wJ5NFV\nF5+ewNxEc/mWvc5u7Ryr7YtvEW962c9QXfD5mONKsnUUsP/nAoIBAERwUmUlL9YP\nq6MSn+bUr6qNZPsVYoQAo8nTjZWiuSjJa2XWnh7sftnISWkwkiiRxo7qfq3sAiD5\nB8+tM6kONeICBXukldXJerxoVBspYa+RiPuDWy2pwGRDBpfty3QqJOpu5g2ThYFJ\nD5Xu0yCuX8ZJRj33nliI8dQgKd0xQQva6p2VuXzyRT8LwXMfRwLuSB6Schc9mF8C\nkWCb4m0ujlEKe1xKoKt2zG9b1o7XyaVhxguSUAuEznifMzsEUfuONJOy+XoQELex\nF0wvLzNzABcyxkM3lx52uG41mKgJiV6Z0ZyuBRvt+V3VL/38tPn9lsTaFi8N6/IH\nRyy0bWP5s44=\n-----END PUBLIC KEY-----\n", "Address": "18.237.147.105", "Tls_certificate": "-----BEGIN CERTIFICATE-----MIIDbDCCAlSgAwIBAgIJAOUNtZneIYECMA0GCSqGSIb3DQEBBQUAMGgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJpcDAeFw0xOTAzMDUxODM1NDNaFw0yOTAzMDIxODM1NDNaMGgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJpcDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPP0WyVkfZA/CEd2DgKpcudn0oDhDwsjmx8LBDWsUgQzyLrFiVigfUmUefknUH3dTJjmiJtGqLsayCnWdqWLHPJYvFfsWYW0IGF93UG/4N5UAWO4okC3CYgKSi4ekpfw2zgZq0gmbzTnXcHF9gfmQ7jJUKSEtJPSNzXq+PZeJTC9zJAb4Lj8QzH18rDM8DaL2y1ns0Y2Hu0edBFn/OqavBJKb/uAm3AEjqeOhC7EQUjVamWlTBPt40+B/6aFJX5BYm2JFkRsGBIyBVL46MvC02MgzTT9bJIJfwqmBaTruwemNgzGu7Jk03hqqS1TUEvSI6/x8bVoba3orcKkf9HsDjECAwEAAaMZMBcwFQYDVR0RBA4wDIIKKi5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEAneUocN4AbcQAC1+b3To8u5UGdaGxhcGyZBlAoenRVdjXK3lTjsMdMWb4QctgNfIfU/zuUn2mxTmF/ekP0gCCgtleZr9+DYKU5hlXk8K10uKxGD6EvoiXZzlfeUuotgp2qvI3ysOm/hvCfyEkqhfHtbxjV7j7v7eQFPbvNaXbLa0yr4C4vMK/Z09Ui9JrZ/Z4cyIkxfC6/rOqAirSdIp09EGiw7GM8guHyggE4IiZrDslT8V3xIl985cbCxSxeW1RtgH4rdEXuVe9+31oJhmXOE9ux2jCop9tEJMgWg7HStrJ5plPbb+HmjoX3nBO04E56m52PyzMNV+2N21IPppKwA==-----END CERTIFICATE-----"}, {"Id": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "Dsa_public_key": "-----BEGIN PUBLIC KEY-----\nMIIDNDCCAiwCggEBAJ22+1lRtmu2/h4UDx0s5VAjdBYf1lON8WSCGGQvC1xIyPek\nGq36GHMkuHZ0+hgisA8ez4E2lD18VXVyZOWhpE/+AS6ZNuAMHT6TELAcfReYBdMF\niyqfS7b5cWv+YRfGtbPMTZvjQRBK1KgK1slOAF9LmT4U8JHrUXQ78zBQw43iNVZ+\nGzTD1qXAzqoaDzaCE8PRmEPQtLCdy5/HLTnI3kHxvxTUu0Vjyig3FiHK0zJLai05\nIUW+v6x0iAUjb1yi/pK4cc2PnDbTKStVCcqMqneirfx7/XfdpvcRJadFb+oVPkMy\nVqImHGoG7TaTeX55lfrVqrvPvj7aJ0HjdUBK4lsCIQDywxGTdM52yTVpkLRlN0oX\n8j+e01CJvZafYcbd6ZmMHwKCAQBcf/awb48UP+gohDNJPkdpxNmIrOW+JaDiSAln\nBxbGE9ewzuaTL4+qfETSyyRSPaU/vk9uw1lYktGqWMQyigbEahVmLn6qcDod7Pi7\nstBdvi65VsFCozhmHRBGHA0TVHIIUFfzSUMJ/6c8YR94syrbtXQMNhyfNb6QmX2y\nAU4u9apheC9Sq+uL1kMsTdCXvFQjsoXa+2DcNk6BYfSio1rKOhCxxNIDzHakcKM6\n/cvdkpWYWavYtW4XJSUteOrGbnG6muPx3SSHGZh0OTzU2DIYaABlR2Dh40wJ5NFV\nF5+ewNxEc/mWvc5u7Ryr7YtvEW962c9QXfD5mONKsnUUsP/nAoIBAFbADcqA8KQh\nxzgylW6VS1dYYelO5DjPZVVSjfdcbj1twu4ZHDNZLOexpv4nGY8xS6vesELXcVOR\n/CHXgh/3byBZYm0zkrBi/FsJJ3nP2uZ1+QCRldI2KzqcLOWH/CAYj8koork9k1Dp\nFq7rMSDgw4pktqvFj9Eev8dSZuRnoCfZbt/6vxi1r30AYAjDYOwcysqcVyUa1tPa\nLEh3JksttXUCd5cvfqatWedTs5Vxo7ICW1toGBHABYvSJkwK0YFfi5RLw+Oda1sA\njJ+aLcIxQjrpoRC2alXCdwmZXVb+O6zluQctw6LJjt4J704ueSvR4VNNhr0uLYGW\nk7e+WoQCS98=\n-----END PUBLIC KEY-----\n", "Address": "52.11.136.238", "Tls_certificate": "-----BEGIN CERTIFICATE-----MIIDbDCCAlSgAwIBAgIJAOUNtZneIYECMA0GCSqGSIb3DQEBBQUAMGgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJpcDAeFw0xOTAzMDUxODM1NDNaFw0yOTAzMDIxODM1NDNaMGgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJpcDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPP0WyVkfZA/CEd2DgKpcudn0oDhDwsjmx8LBDWsUgQzyLrFiVigfUmUefknUH3dTJjmiJtGqLsayCnWdqWLHPJYvFfsWYW0IGF93UG/4N5UAWO4okC3CYgKSi4ekpfw2zgZq0gmbzTnXcHF9gfmQ7jJUKSEtJPSNzXq+PZeJTC9zJAb4Lj8QzH18rDM8DaL2y1ns0Y2Hu0edBFn/OqavBJKb/uAm3AEjqeOhC7EQUjVamWlTBPt40+B/6aFJX5BYm2JFkRsGBIyBVL46MvC02MgzTT9bJIJfwqmBaTruwemNgzGu7Jk03hqqS1TUEvSI6/x8bVoba3orcKkf9HsDjECAwEAAaMZMBcwFQYDVR0RBA4wDIIKKi5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEAneUocN4AbcQAC1+b3To8u5UGdaGxhcGyZBlAoenRVdjXK3lTjsMdMWb4QctgNfIfU/zuUn2mxTmF/ekP0gCCgtleZr9+DYKU5hlXk8K10uKxGD6EvoiXZzlfeUuotgp2qvI3ysOm/hvCfyEkqhfHtbxjV7j7v7eQFPbvNaXbLa0yr4C4vMK/Z09Ui9JrZ/Z4cyIkxfC6/rOqAirSdIp09EGiw7GM8guHyggE4IiZrDslT8V3xIl985cbCxSxeW1RtgH4rdEXuVe9+31oJhmXOE9ux2jCop9tEJMgWg7HStrJ5plPbb+HmjoX3nBO04E56m52PyzMNV+2N21IPppKwA==-----END CERTIFICATE-----"}, {"Id": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "Dsa_public_key": "-----BEGIN PUBLIC KEY-----\nMIIDNTCCAiwCggEBAJ22+1lRtmu2/h4UDx0s5VAjdBYf1lON8WSCGGQvC1xIyPek\nGq36GHMkuHZ0+hgisA8ez4E2lD18VXVyZOWhpE/+AS6ZNuAMHT6TELAcfReYBdMF\niyqfS7b5cWv+YRfGtbPMTZvjQRBK1KgK1slOAF9LmT4U8JHrUXQ78zBQw43iNVZ+\nGzTD1qXAzqoaDzaCE8PRmEPQtLCdy5/HLTnI3kHxvxTUu0Vjyig3FiHK0zJLai05\nIUW+v6x0iAUjb1yi/pK4cc2PnDbTKStVCcqMqneirfx7/XfdpvcRJadFb+oVPkMy\nVqImHGoG7TaTeX55lfrVqrvPvj7aJ0HjdUBK4lsCIQDywxGTdM52yTVpkLRlN0oX\n8j+e01CJvZafYcbd6ZmMHwKCAQBcf/awb48UP+gohDNJPkdpxNmIrOW+JaDiSAln\nBxbGE9ewzuaTL4+qfETSyyRSPaU/vk9uw1lYktGqWMQyigbEahVmLn6qcDod7Pi7\nstBdvi65VsFCozhmHRBGHA0TVHIIUFfzSUMJ/6c8YR94syrbtXQMNhyfNb6QmX2y\nAU4u9apheC9Sq+uL1kMsTdCXvFQjsoXa+2DcNk6BYfSio1rKOhCxxNIDzHakcKM6\n/cvdkpWYWavYtW4XJSUteOrGbnG6muPx3SSHGZh0OTzU2DIYaABlR2Dh40wJ5NFV\nF5+ewNxEc/mWvc5u7Ryr7YtvEW962c9QXfD5mONKsnUUsP/nAoIBAQCN19tTnkS3\nitBQXXR/h8OKl+rliFBLgO6h6GvZL4yQDZFtBAOmkrs3wLoDroJRGCeqz/IUb+JF\njslEr/mpm2kcmK77hr535dq7HsWz1fFl9YyGTaOH055FLSV9QEPAV9j3zWADdQ1v\nuSQll+QfWi6lIibWV4HNQ2ywRFoOY8OBLCJB90UXLeJpaPanpqiM8hjda2VGRDbi\nIixEE2lCOWITydiz2DmvXrLhVGF49+g5MDwbWO65dmasCe//Ff6Z4bJ6n040x9xv\nVtac8nX6FO3eBsV5d+rG6HZXSG3brCKRCSKYCTX1IkTSiutYxYqvwaluoCjOakh0\nKkqvQ8IeVZ+B\n-----END PUBLIC KEY-----\n", "Address": "34.213.79.31", "Tls_certificate": "-----BEGIN CERTIFICATE-----MIIDbDCCAlSgAwIBAgIJAOUNtZneIYECMA0GCSqGSIb3DQEBBQUAMGgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJpcDAeFw0xOTAzMDUxODM1NDNaFw0yOTAzMDIxODM1NDNaMGgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJpcDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPP0WyVkfZA/CEd2DgKpcudn0oDhDwsjmx8LBDWsUgQzyLrFiVigfUmUefknUH3dTJjmiJtGqLsayCnWdqWLHPJYvFfsWYW0IGF93UG/4N5UAWO4okC3CYgKSi4ekpfw2zgZq0gmbzTnXcHF9gfmQ7jJUKSEtJPSNzXq+PZeJTC9zJAb4Lj8QzH18rDM8DaL2y1ns0Y2Hu0edBFn/OqavBJKb/uAm3AEjqeOhC7EQUjVamWlTBPt40+B/6aFJX5BYm2JFkRsGBIyBVL46MvC02MgzTT9bJIJfwqmBaTruwemNgzGu7Jk03hqqS1TUEvSI6/x8bVoba3orcKkf9HsDjECAwEAAaMZMBcwFQYDVR0RBA4wDIIKKi5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEAneUocN4AbcQAC1+b3To8u5UGdaGxhcGyZBlAoenRVdjXK3lTjsMdMWb4QctgNfIfU/zuUn2mxTmF/ekP0gCCgtleZr9+DYKU5hlXk8K10uKxGD6EvoiXZzlfeUuotgp2qvI3ysOm/hvCfyEkqhfHtbxjV7j7v7eQFPbvNaXbLa0yr4C4vMK/Z09Ui9JrZ/Z4cyIkxfC6/rOqAirSdIp09EGiw7GM8guHyggE4IiZrDslT8V3xIl985cbCxSxeW1RtgH4rdEXuVe9+31oJhmXOE9ux2jCop9tEJMgWg7HStrJ5plPbb+HmjoX3nBO04E56m52PyzMNV+2N21IPppKwA==-----END CERTIFICATE-----"}], "registration": {"Dsa_public_key": "-----BEGIN PUBLIC KEY-----\nMIIDNDCCAiwCggEBAJ22+1lRtmu2/h4UDx0s5VAjdBYf1lON8WSCGGQvC1xIyPek\nGq36GHMkuHZ0+hgisA8ez4E2lD18VXVyZOWhpE/+AS6ZNuAMHT6TELAcfReYBdMF\niyqfS7b5cWv+YRfGtbPMTZvjQRBK1KgK1slOAF9LmT4U8JHrUXQ78zBQw43iNVZ+\nGzTD1qXAzqoaDzaCE8PRmEPQtLCdy5/HLTnI3kHxvxTUu0Vjyig3FiHK0zJLai05\nIUW+v6x0iAUjb1yi/pK4cc2PnDbTKStVCcqMqneirfx7/XfdpvcRJadFb+oVPkMy\nVqImHGoG7TaTeX55lfrVqrvPvj7aJ0HjdUBK4lsCIQDywxGTdM52yTVpkLRlN0oX\n8j+e01CJvZafYcbd6ZmMHwKCAQBcf/awb48UP+gohDNJPkdpxNmIrOW+JaDiSAln\nBxbGE9ewzuaTL4+qfETSyyRSPaU/vk9uw1lYktGqWMQyigbEahVmLn6qcDod7Pi7\nstBdvi65VsFCozhmHRBGHA0TVHIIUFfzSUMJ/6c8YR94syrbtXQMNhyfNb6QmX2y\nAU4u9apheC9Sq+uL1kMsTdCXvFQjsoXa+2DcNk6BYfSio1rKOhCxxNIDzHakcKM6\n/cvdkpWYWavYtW4XJSUteOrGbnG6muPx3SSHGZh0OTzU2DIYaABlR2Dh40wJ5NFV\nF5+ewNxEc/mWvc5u7Ryr7YtvEW962c9QXfD5mONKsnUUsP/nAoIBAAlELnrXLG0s\n4yAAn7IsVWwY7swDnbBcsIF2cnef4tjm/nNwrFKp5AxYqgeXCiJM8VkyJrotWG50\nnXQwMCR6BsvYrlVt/RmQvR8BSrir62uSLK7hMKm7dXnFvtyFtjp91UwTRbIjxhUQ\nGYnhAzrkCDWo1m54ysqXEGlrVwvRXrCAXiLKPiTEIS+B4GFH9W26SwBxhFLNYSUk\nZZ7+4qwMf9aTu7kIpXTP3hNIyRCjtuZvo5SnymtbLARwTP943hW8MOj0+Ege+m1P\ntey6rkMUGQ86cgK9/7+7Jb+EwW5UxdQtFPUFeNKdQ6zDPS6qbliecUrsc12tdgeg\nhQyuMbyKUuo=\n-----END PUBLIC KEY-----\n", "Address": "registration.default.cmix.rip"}, "udb": {"Id": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3], "Dsa_public_key": "-----BEGIN PUBLIC KEY-----\nMIIDNDCCAiwCggEBAJ22+1lRtmu2/h4UDx0s5VAjdBYf1lON8WSCGGQvC1xIyPek\nGq36GHMkuHZ0+hgisA8ez4E2lD18VXVyZOWhpE/+AS6ZNuAMHT6TELAcfReYBdMF\niyqfS7b5cWv+YRfGtbPMTZvjQRBK1KgK1slOAF9LmT4U8JHrUXQ78zBQw43iNVZ+\nGzTD1qXAzqoaDzaCE8PRmEPQtLCdy5/HLTnI3kHxvxTUu0Vjyig3FiHK0zJLai05\nIUW+v6x0iAUjb1yi/pK4cc2PnDbTKStVCcqMqneirfx7/XfdpvcRJadFb+oVPkMy\nVqImHGoG7TaTeX55lfrVqrvPvj7aJ0HjdUBK4lsCIQDywxGTdM52yTVpkLRlN0oX\n8j+e01CJvZafYcbd6ZmMHwKCAQBcf/awb48UP+gohDNJPkdpxNmIrOW+JaDiSAln\nBxbGE9ewzuaTL4+qfETSyyRSPaU/vk9uw1lYktGqWMQyigbEahVmLn6qcDod7Pi7\nstBdvi65VsFCozhmHRBGHA0TVHIIUFfzSUMJ/6c8YR94syrbtXQMNhyfNb6QmX2y\nAU4u9apheC9Sq+uL1kMsTdCXvFQjsoXa+2DcNk6BYfSio1rKOhCxxNIDzHakcKM6\n/cvdkpWYWavYtW4XJSUteOrGbnG6muPx3SSHGZh0OTzU2DIYaABlR2Dh40wJ5NFV\nF5+ewNxEc/mWvc5u7Ryr7YtvEW962c9QXfD5mONKsnUUsP/nAoIBACvR2lUslz3D\nB/MUo0rHVIHVkhVJCxNjtgTOYgJ9ckArSXQbYzr/fcigcNGjUO2LbK5NFp9GK43C\nrLxMUnJ9nkyIVPaWvquJFZItjcDK3NiNGyD4XyM0eRj4dYeSxQM48hvFbmtbjlXn\n9SQTnGIlr1XnTI4RVHZSQOL6kFJIaLw6wYrQ4w08Ng+p45brp5ercAHnLiftNUWP\nqROhQkdSEpS9LEwfotUSY1jP2AhQfaIMxaeXsZuTU1IYvdhMFRL3DR0r5Ww2Upf8\ng0Ace0mtnsUQ2OG+7MTh2jYIEWRjvuoe3RCz603ujW6g7BfQ1H7f4YFwc5xOOJ3u\nr4dj49dCCjc=\n-----END PUBLIC KEY-----\n"}, "E2e": {"Prime": "0xFFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF", "Small_prime": "0x7FFFFFFFFFFFFFFFE487ED5110B4611A62633145C06E0E68948127044533E63A0105DF531D89CD9128A5043CC71A026EF7CA8CD9E69D218D98158536F92F8A1BA7F09AB6B6A8E122F242DABB312F3F637A262174D31BF6B585FFAE5B7A035BF6F71C35FDAD44CFD2D74F9208BE258FF324943328F6722D9EE1003E5C50B1DF82CC6D241B0E2AE9CD348B1FD47E9267AFC1B2AE91EE51D6CB0E3179AB1042A95DCF6A9483B84B4B36B3861AA7255E4C0278BA3604650C10BE19482F23171B671DF1CF3B960C074301CD93C1D17603D147DAE2AEF837A62964EF15E5FB4AAC0B8C1CCAA4BE754AB5728AE9130C4C7D02880AB9472D455655347FFFFFFFFFFFFFFF", "Generator": "0x02"}, "Cmix": {"Prime": "0xFFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF", "Small_prime": "0x7FFFFFFFFFFFFFFFE487ED5110B4611A62633145C06E0E68948127044533E63A0105DF531D89CD9128A5043CC71A026EF7CA8CD9E69D218D98158536F92F8A1BA7F09AB6B6A8E122F242DABB312F3F637A262174D31BF6B585FFAE5B7A035BF6F71C35FDAD44CFD2D74F9208BE258FF324943328F6722D9EE1003E5C50B1DF82CC6D241B0E2AE9CD348B1FD47E9267AFC1B2AE91EE51D6CB0E3179AB1042A95DCF6A9483B84B4B36B3861AA7255E4C0278BA3604650C10BE19482F23171B671DF1CF3B960C074301CD93C1D17603D147DAE2AEF837A62964EF15E5FB4AAC0B8C1CCAA4BE754AB5728AE9130C4C7D02880AB9472D455655347FFFFFFFFFFFFFFF", "Generator": "0x02"}}
MEQCICf7Iivf4IKOsmCqKmT6PfyiG3URH6kD/+8ngzze4cA4AiA9P4kQ0nUVaHZobawrJ1YpVjSTxTkdPOp6FLf1qNBdaQ==`

func TestDecodeNDF(t *testing.T) {
	jsonNDF, err := DecodeNDF(ExampleNDF2)
	if err != nil {
		t.Error(err)
	}
	// Ensure fields are populated
	if jsonNDF.Timestamp == "" {
        t.Error("Timestamp wasn't populated")
	}
	if jsonNDF.Nodes == nil {
		t.Fatal("Nodes weren't populated")
	}
	if jsonNDF.Gateways == nil {
		t.Fatal("Gateways weren't populated")
	}
	if jsonNDF.Cmix.Generator == "" {
		t.Error("Cmix Generator wasn't populated")
	}
	if jsonNDF.Cmix.Prime == "" {
		t.Error("Cmix Prime wasn't populated")
	}
	if jsonNDF.Cmix.Small_prime == "" {
		t.Error("Cmix Small Prime wasn't populated")
	}
	if jsonNDF.E2e.Generator == "" {
		t.Error("e2e generator wasn't populated")
	}
	if jsonNDF.E2e.Prime == "" {
		t.Error("e2e prime wasn't populated")
	}
	if jsonNDF.E2e.Small_prime == "" {
		t.Error("e2e small prime wasn't populated")
	}
	if jsonNDF.Registration.Address == "" {
		t.Error("Registration address wasn't populated")
	}
	if jsonNDF.Registration.Dsa_public_key == "" {
		t.Error("Registration publickey wasn't populated")
	}
	if jsonNDF.Gateways[0].Address == "" {
		t.Error("Gateway address wasn't populated")
	}
	if jsonNDF.Gateways[0].Tls_certificate == "" {
		t.Error("Gateway TLS wasn't populated")
	}
	if jsonNDF.Nodes[0].Tls_certificate == "" {
		t.Error("Node TLS wasn't populated")
	}
	if jsonNDF.Nodes[0].Address == "" {
		t.Error("Node Address wasn't populated")
	}
	if jsonNDF.Nodes[0].Dsa_public_key == "" {
		t.Error("Node publickey wasn't populated")
	}
	if jsonNDF.Nodes[0].Id == nil {
		t.Error("Node ID wasn't populated")
	}
	if jsonNDF.Udb.Id == nil {
		t.Error("UDB ID wasn't populated")
	}
	if jsonNDF.Udb.Dsa_public_key == "" {
		t.Error("UDB publickey wasn't populated")
	}
}
