package site

import(
    "github.com/kenan-rhoton/vineyard/site/admin"
)

func Setup() {
    admin.Setup()
    ResourcesSetup()
    LandingSetup()
    LoginSetup()
}
