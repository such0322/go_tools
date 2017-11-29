package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"odin_tools/models/bonus"
	BMi "odin_tools/models/bridge/misc"
	"strconv"
)

type GiftController struct {
	Controller
}

func (c GiftController) List(w http.ResponseWriter, r *http.Request) {
	var p int
	if r.FormValue("p") == "" {
		p = 1
	} else {
		var err error
		p, err = strconv.Atoi(r.FormValue("p"))
		if err != nil {
			p = 1
		}
	}

	giftCodes := BMi.GiftCodes{}
	_, pager := giftCodes.GetPage(p, 100, "/gift/list")

	c.data = make(map[string]interface{})
	c.data["CodeList"] = giftCodes.Data
	c.data["Pager"] = template.HTML(pager)
	c.data["MapType"] = BMi.GiftCode.GetTypeMap
	c.Render(w, r)
}

func (c GiftController) NewGift(w http.ResponseWriter, r *http.Request) {
	c.data = make(map[string]interface{})
	c.data["RewardType"] = bonus.RewardType
	c.Render(w, r)
}

func (c GiftController) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

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
	reward_types := r.Form["reward_type"]
	reward_ids := r.Form["reward_id"]
	reward_qs := r.Form["reward_quantity"]
	bd := bonus.BonusData{}
	bds := bd.CreateRewards(reward_types, reward_ids, reward_qs)
	data["package"] = bds
	data["startDate"] = r.FormValue("startDate")
	data["endDate"] = r.FormValue("endDate")

	BMi.CreateGiftCodes(data)

	// http.Redirect(w, r, "/gift/list", http.StatusFound)
}

func (c GiftController) RandomCode(w http.ResponseWriter, r *http.Request) {
	code := BMi.GetRandomCode()
	fmt.Fprintf(w, code)
}

func (c GiftController) GetBounsAll(w http.ResponseWriter, r *http.Request) {
	rewardType := r.FormValue("reward_type")
	bonus, err := bonus.NewBonus(rewardType)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	bonus.GetRewardNames()
	fmt.Fprint(w, bonus.ToJson())
}
