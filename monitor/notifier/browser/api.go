package notifier

import (
	"net/http"

	goutils "github.com/anonymous-author-00000/go-utils"
	"github.com/anonymous-author-00000/ra-webs/monitor"
	"github.com/labstack/echo/v4"
)

var postSubscribeApi = goutils.EchoRoute[monitor.Monitor]{
	Method: goutils.POST,
	Path:   "/subscription",
	F: func(monitor *monitor.Monitor) goutils.EchoRouteFunc {
		return func(c echo.Context) error {
			var data struct {
				Subscription struct {
					Endpoint       string `json:"endpoint"`
					ExpirationTime int    `json:"expirationTime"`
					Keys           struct {
						Auth   string `json:"auth"`
						P256dh string `json:"p256dh"`
					} `json:"keys"`
				} `json:"subscription"`
			}

			if err := c.Bind(&data); err != nil {
				return err
			}

			subscription, err := monitor.DB.Client.Subscription.
				Create().
				SetEndpoint(data.Subscription.Endpoint).
				SetAuth(data.Subscription.Keys.Auth).
				SetP256dh(data.Subscription.Keys.P256dh).
				Save(*monitor.DB.Ctx)

			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, subscription)
		}
	},
}

func getSubscriptionConfigApi(notifier *BrowserNotifier) goutils.EchoRoute[monitor.Monitor] {
	return goutils.EchoRoute[monitor.Monitor]{
		Method: goutils.GET,
		Path:   "/config/subscription",
		F: func(_ *monitor.Monitor) goutils.EchoRouteFunc {
			return func(c echo.Context) error {
				VapidPublicKey := notifier.VapidPublicKey

				var data struct {
					VapidPublicKey string `json:"vapid_public_key"`
				}

				data.VapidPublicKey = VapidPublicKey

				return c.JSON(http.StatusOK, data)
			}
		},
	}
}

func (notifier *BrowserNotifier) setUpRoute(e *echo.Group) {
	postSubscribeApi.Set(e, notifier.Monitor)
	getSubscriptionConfigApi(notifier).Set(e, notifier.Monitor)
}
