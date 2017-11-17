package handler

import (
	"fmt"
	"net/http"
	"odin_tools/models/bonus"
	BMi "odin_tools/models/bridge/misc"
	"strconv"
)

type GiftController struct {
	Controller
}

func (c GiftController) List(w http.ResponseWriter, r *http.Request) {
	giftCodes := BMi.GiftCodes{}
	giftCodes.GetAll()

	c.data = make(map[string]interface{})
	c.data["CodeList"] = giftCodes.Data

	c.tpl = "gift/list"
	c.Render(w, r)
}

func (c GiftController) NewGift(w http.ResponseWriter, r *http.Request) {
	c.data = make(map[string]interface{})
	c.data["RewardType"] = bonus.RewardType
	c.tpl = "gift/new"
	c.Render(w, r)
}

func (c GiftController) Create(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["code"] = r.FormValue("code")
	data["type"], _ = strconv.Atoi(r.FormValue("type"))
	data["channel"] = r.FormValue("channel")
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	if quantity <= 0 {
		//TODO 做个flash，跳转后弹出错误信息
		http.Redirect(w, r, "/gift/new", http.StatusFound)
	}
	data["quantity"] = quantity
	data["startDate"] = r.FormValue("startDate")
	data["endDate"] = r.FormValue("endDate")
	BMi.CreateGiftCodes(data)

	http.Redirect(w, r, "/gift/list", http.StatusFound)
}

func (c GiftController) RandomCode(w http.ResponseWriter, r *http.Request) {
	code := BMi.GetRandomCode()
	fmt.Fprintf(w, code)
}

func (c GiftController) GetBounsAll(w http.ResponseWriter, r *http.Request) {
	rewardType := r.FormValue("reward_type")
	switch rewardType {
	case bonus.TypeItem:
		bi := bonus.BounsItems{}
		bi.GetAll()
	}

	fmt.Fprintf(w, rewardType)
}
