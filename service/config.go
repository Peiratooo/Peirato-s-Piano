package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var configFilePath = "config.json"

type Color struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Window struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

type Config struct {
	Colors       map[string]Color `json:"colors"`
	KeyLabel     string           `json:"keyLabel"`
	KeyboardType int              `json:"keyboardType"`
	Velocity     uint8            `json:"velocity"`
	Opacity      int              `json:"opacity"`
	Version      string           `json:"version"`
	ShowPedal    bool             `json:"showPedal"`
	Volume       int32            `json:"volume"`
}

var DefaultConfig = Config{
	Colors: map[string]Color{
		"whiteKey": {
			Label: "白键按下",
			Color: "#9AF7B3",
		},
		"blackKey": {
			Label: "黑键按下",
			Color: "#5FFF5F",
		},
		"damperPedal": {
			Label: "延音踏板踩下",
			Color: "#e7b510",
		},
		"softPedal": {
			Label: "柔音踏板踩下",
			Color: "#10e786",
		},
		"sostenutoPedal": {
			Label: "消音踏板踩下",
			Color: "#1054e7",
		},
	},
	KeyLabel:     "octave_key",
	KeyboardType: 0,
	Velocity:     80,
	Volume:       80,
	Opacity:      100,
	ShowPedal:    true,
}

var UserConfig Config

func LoadConfig(version string) {
	ucd, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("获取用户配置目录失败:", err)
		ucd = "./assets"
	}
	configFilePath = filepath.Join(ucd, "Peirato's Piano", "config.json")
	var config Config
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {

		err := SaveConfig(DefaultConfig)
		if err != nil {
			fmt.Println("保存默认配置失败:", err)
		}
		config = DefaultConfig
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("读取配置文件失败:", err)
		config = DefaultConfig
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("解析配置文件失败:", err)
		config = DefaultConfig
	}

	if config.Opacity < 20 {
		config.Opacity = 100
		err = SaveConfig(config)
		fmt.Println(err)
	}
	if config.Version != version {
		config.Version = version
		err = SaveConfig(config)
		fmt.Println(err)
	}
	UserConfig = config
}

func SaveConfig(config Config) error {

	UserConfig = config

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化 JSON 失败: %v", err)
	}

	dir := filepath.Dir(configFilePath)
	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("创建目录失败: %v", err)
		}
	}

	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	return nil
}

func (k *Keyboard) SendConfig() Config {
	return UserConfig
}

func (k *Keyboard) ReceiveConfig(config Config) (bool, string) {
	if err := SaveConfig(config); err != nil {
		return false, err.Error()
	} else {
		return true, ""
	}
}

func (k *Keyboard) ResetConfig() Config {
	SaveConfig(DefaultConfig)
	return DefaultConfig
}
