package main

import (
	"context"
	"embed"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/slider"
)

type slideData struct {
	Text  string
	Image string
}

var slidesData = []slideData{
	{Image: "youtube.png", Text: "*0\\. YouTube* is an American online video sharing and social media platform headquartered in San Bruno, California\\. It was launched on _February 14, 2005_, by *Steve Chen*, *Chad Hurley*, and *Jawed Karim*"},
	{Image: "vk.png", Text: "*1\\. VK* \\(short for its original name VKontakte; Russian: ВКонтакте, meaning InContact\\) is a Russian online social media and social networking service based in *Saint Petersburg*"},
	{Image: "skype.png", Text: "*2\\. Skype* is a proprietary telecommunications application operated by Skype Technologies, a division of *Microsoft*, best known for VoIP\\-based videotelephony, videoconferencing and voice calls"},
	{Image: "reddit.png", Text: "*3\\. Reddit* \\(\\/ˈrɛdɪt\\/, stylized as reddit\\) is an American social news aggregation, web content rating, and discussion website"},
	{Image: "twitter.png", Text: "*4\\. Twitter* is an American microblogging and social networking service on which users post and interact with messages known as *tweets*"},
	{Image: "pinterest.png", Text: "*5\\. Pinterest* is an image sharing and social media service designed to enable saving and discovery of information on the internet using images, and on a smaller scale, animated GIFs and videos, in the form of pinboards"},
	{Image: "instagram.png", Text: "*6\\. Instagram* is an American photo and video sharing social networking service founded by *Kevin Systrom* and *Mike Krieger*\\. In April 2012, Facebook Inc\\. acquired the service for approximately *US$1 billion* in cash and stock"},
	{Image: "linkedin.png", Text: "*7\\. LinkedIn* is an American business and employment\\-oriented online service that operates via websites and mobile apps\\. Launched on May 5, 2003"},
	{Image: "facebook.png", Text: "*8\\. Facebook* is an American online social media and social networking service owned by Meta Platforms\\. Founded in 2004 by *Mark Zuckerberg* with fellow Harvard College students and roommates *Eduardo Saverin*, *Andrew McCollum*, *Dustin Moskovitz*, and *Chris Hughes*"},
}

//go:embed images
var sliderImages embed.FS

func handlerSlider(ctx context.Context, b *bot.Bot, update *models.Update) {
	slides := generateSlides()

	opts := []slider.Option{
		slider.OnSelect("Select", true, sliderOnSelect),
		slider.OnCancel("Cancel", true, sliderOnCancel),
	}

	sl := slider.New(slides, opts...)

	sl.Show(ctx, b, update.Message.Chat.ID)
}

func sliderOnSelect(ctx context.Context, b *bot.Bot, message *models.Message, item int) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: message.Chat.ID,
		Text:   "Select " + strconv.Itoa(item),
	})
}

func sliderOnCancel(ctx context.Context, b *bot.Bot, message *models.Message) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: message.Chat.ID,
		Text:   "Cancel",
	})
}

func generateSlides() []slider.Slide {
	var slides []slider.Slide

	for _, v := range slidesData {

		imageContent, _ := sliderImages.ReadFile("images/" + v.Image)

		slides = append(slides, slider.Slide{
			Photo:    string(imageContent),
			IsUpload: true,
			Text:     v.Text,
		})
	}

	return slides
}
