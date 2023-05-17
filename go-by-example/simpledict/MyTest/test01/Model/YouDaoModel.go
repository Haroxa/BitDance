package Model

type YouDaoResponse struct {
	WebTrans struct {
		WebTranslation []struct {
			Same      string `json:"@same,omitempty"`
			Key       string `json:"key"`
			KeySpeech string `json:"key-speech"`
			Trans     []struct {
				Summary struct {
					Line []string `json:"line"`
				} `json:"summary"`
				Value   string `json:"value"`
				Support int    `json:"support"`
				URL     string `json:"url"`
			} `json:"trans"`
		} `json:"web-translation"`
	} `json:"web_trans"`
	OxfordAdvanceHTML struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxfordAdvanceHtml"`
	PicDict struct {
		Pic []struct {
			Image string `json:"image"`
			Host  string `json:"host"`
			URL   string `json:"url"`
		} `json:"pic"`
	} `json:"pic_dict"`
	VideoSents struct {
		SentsData []struct {
			VideoCover  string `json:"video_cover"`
			Contributor string `json:"contributor"`
			SubtitleSrt string `json:"subtitle_srt"`
			ID          int    `json:"id"`
			Video       string `json:"video"`
		} `json:"sents_data"`
		WordInfo struct {
			ReturnPhrase string   `json:"return-phrase"`
			Sense        []string `json:"sense"`
		} `json:"word_info"`
	} `json:"video_sents"`
	Simple struct {
		Query string `json:"query"`
		Word  []struct {
			Usphone      string `json:"usphone"`
			Ukphone      string `json:"ukphone"`
			Ukspeech     string `json:"ukspeech"`
			ReturnPhrase string `json:"return-phrase"`
			MultiPhone   struct {
				Uk []struct {
					Phone  string   `json:"phone"`
					Pos    []string `json:"pos"`
					Speech string   `json:"speech"`
				} `json:"uk"`
				Us []struct {
					Phone  string   `json:"phone"`
					Pos    []string `json:"pos"`
					Speech string   `json:"speech"`
				} `json:"us"`
			} `json:"multiPhone"`
			Usspeech string `json:"usspeech"`
		} `json:"word"`
	} `json:"simple"`
	Phrs struct {
		Word string `json:"word"`
		Phrs []struct {
			Headword    string `json:"headword"`
			Translation string `json:"translation"`
			Source      string `json:"source,omitempty"`
		} `json:"phrs"`
	} `json:"phrs"`
	Oxford struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxford"`
	Syno struct {
		Synos []struct {
			Pos  string   `json:"pos"`
			Ws   []string `json:"ws"`
			Tran string   `json:"tran"`
		} `json:"synos"`
		Word string `json:"word"`
	} `json:"syno"`
	Collins struct {
		SuperHeadwords struct {
			SuperHeadword []string `json:"super_headword"`
		} `json:"super_headwords"`
		CollinsEntries []struct {
			SuperHeadword string `json:"super_headword,omitempty"`
			Entries       struct {
				Entry []struct {
					TranEntry []struct {
						PosEntry struct {
							Pos     string `json:"pos"`
							PosTips string `json:"pos_tips"`
						} `json:"pos_entry"`
						ExamSents struct {
							Sent []struct {
								ChnSent string `json:"chn_sent"`
								EngSent string `json:"eng_sent"`
							} `json:"sent"`
						} `json:"exam_sents"`
						Tran string `json:"tran"`
					} `json:"tran_entry"`
				} `json:"entry"`
			} `json:"entries,omitempty"`
			Phonetic     string `json:"phonetic,omitempty"`
			BasicEntries struct {
				BasicEntry []struct {
					Cet       string `json:"cet"`
					Headword  string `json:"headword"`
					Wordforms struct {
						Wordform []struct {
							Word string `json:"word"`
						} `json:"wordform"`
					} `json:"wordforms"`
				} `json:"basic_entry"`
			} `json:"basic_entries"`
			Headword string `json:"headword"`
			Star     string `json:"star,omitempty"`
		} `json:"collins_entries"`
	} `json:"collins"`
	WordVideo struct {
		WordVideos []struct {
			Ad struct {
				Avatar string `json:"avatar"`
				Title  string `json:"title"`
				URL    string `json:"url"`
			} `json:"ad"`
			Video struct {
				Cover string `json:"cover"`
				Image string `json:"image"`
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"video"`
		} `json:"word_videos"`
	} `json:"word_video"`
	Webster struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"webster"`
	Discriminate struct {
		Data []struct {
			Source string `json:"source"`
			Usages []struct {
				Headword string `json:"headword"`
				Usage    string `json:"usage"`
			} `json:"usages"`
			Headwords []string `json:"headwords"`
			Tran      string   `json:"tran"`
		} `json:"data"`
		ReturnPhrase string `json:"return-phrase"`
	} `json:"discriminate"`
	WikipediaDigest struct {
		Summarys []struct {
			Summary string `json:"summary"`
			Key     string `json:"key"`
		} `json:"summarys"`
		Source struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"source"`
	} `json:"wikipedia_digest"`
	Lang string `json:"lang"`
	Ec   struct {
		WebTrans []string `json:"web_trans"`
		Special  []struct {
			Nat   string `json:"nat"`
			Major string `json:"major"`
		} `json:"special"`
		ExamType []string `json:"exam_type"`
		Source   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"source"`
		Word struct {
			Usphone  string `json:"usphone"`
			Ukphone  string `json:"ukphone"`
			Ukspeech string `json:"ukspeech"`
			Trs      []struct {
				Pos  string `json:"pos"`
				Tran string `json:"tran"`
			} `json:"trs"`
			Wfs []struct {
				Wf struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"wf"`
			} `json:"wfs"`
			ReturnPhrase string `json:"return-phrase"`
			Usspeech     string `json:"usspeech"`
		} `json:"word"`
	} `json:"ec"`
	Ee struct {
		Source struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"source"`
		Word struct {
			Trs []struct {
				Pos string `json:"pos"`
				Tr  []struct {
					Tran         string   `json:"tran"`
					SimilarWords []string `json:"similar-words"`
				} `json:"tr"`
			} `json:"trs"`
			Phone        string `json:"phone"`
			Speech       string `json:"speech"`
			ReturnPhrase string `json:"return-phrase"`
		} `json:"word"`
	} `json:"ee"`
	BlngSentsPart struct {
		SentenceCount int `json:"sentence-count"`
		SentencePair  []struct {
			Sentence            string `json:"sentence"`
			SentenceEng         string `json:"sentence-eng"`
			SentenceTranslation string `json:"sentence-translation"`
			SpeechSize          string `json:"speech-size"`
			AlignedWords        struct {
				Src struct {
					Chars []struct {
						S      string `json:"@s"`
						E      string `json:"@e"`
						Aligns struct {
							Sc []struct {
								ID string `json:"@id"`
							} `json:"sc"`
							Tc []struct {
								ID string `json:"@id"`
							} `json:"tc"`
						} `json:"aligns"`
						ID string `json:"@id"`
					} `json:"chars"`
				} `json:"src"`
				Tran struct {
					Chars []struct {
						S      string `json:"@s"`
						E      string `json:"@e"`
						Aligns struct {
							Sc []struct {
								ID string `json:"@id"`
							} `json:"sc"`
							Tc []struct {
								ID string `json:"@id"`
							} `json:"tc"`
						} `json:"aligns"`
						ID string `json:"@id"`
					} `json:"chars"`
				} `json:"tran"`
			} `json:"aligned-words"`
			Source         string `json:"source"`
			URL            string `json:"url"`
			SentenceSpeech string `json:"sentence-speech"`
		} `json:"sentence-pair"`
		More        string `json:"more"`
		TrsClassify []struct {
			Proportion string `json:"proportion"`
			Tr         string `json:"tr"`
		} `json:"trs-classify"`
	} `json:"blng_sents_part"`
	Individual struct {
		Trs []struct {
			Pos  string `json:"pos"`
			Tran string `json:"tran"`
		} `json:"trs"`
		Idiomatic []struct {
			Colloc struct {
				En string `json:"en"`
				Zh string `json:"zh"`
			} `json:"colloc"`
		} `json:"idiomatic"`
		Level    string `json:"level"`
		ExamInfo struct {
			Year             int `json:"year"`
			QuestionTypeInfo []struct {
				Time int    `json:"time"`
				Type string `json:"type"`
			} `json:"questionTypeInfo"`
			RecommendationRate int `json:"recommendationRate"`
			Frequency          int `json:"frequency"`
		} `json:"examInfo"`
		ReturnPhrase  string `json:"return-phrase"`
		PastExamSents []struct {
			En     string `json:"en"`
			Source string `json:"source"`
			Zh     string `json:"zh"`
		} `json:"pastExamSents"`
	} `json:"individual"`
	CollinsPrimary struct {
		Words struct {
			Indexforms []string `json:"indexforms"`
			Word       string   `json:"word"`
		} `json:"words"`
		Gramcat []struct {
			Audiourl      string `json:"audiourl"`
			Pronunciation string `json:"pronunciation"`
			Senses        []struct {
				Sensenumber string `json:"sensenumber"`
				Examples    []struct {
					Sense struct {
						Lang string `json:"lang"`
						Word string `json:"word"`
					} `json:"sense"`
					Example string `json:"example"`
				} `json:"examples"`
				Definition string `json:"definition"`
				Lang       string `json:"lang"`
				Word       string `json:"word"`
			} `json:"senses"`
			Partofspeech string `json:"partofspeech"`
			Audio        string `json:"audio"`
			Forms        []struct {
				Form string `json:"form"`
			} `json:"forms"`
			Phrases []struct {
				Phrase string `json:"phrase,omitempty"`
				Senses []struct {
					Examples []struct {
						Sense struct {
							Lang string `json:"lang"`
							Word string `json:"word"`
						} `json:"sense"`
						Example string `json:"example"`
					} `json:"examples"`
					Definition string `json:"definition"`
					Lang       string `json:"lang"`
					Word       string `json:"word"`
				} `json:"senses"`
				Phrasalverb []struct {
					Phrase string `json:"phrase"`
				} `json:"phrasalverb,omitempty"`
			} `json:"phrases,omitempty"`
		} `json:"gramcat"`
	} `json:"collins_primary"`
	RelWord struct {
		Word string `json:"word"`
		Stem string `json:"stem"`
		Rels []struct {
			Rel struct {
				Pos   string `json:"pos"`
				Words []struct {
					Word string `json:"word"`
					Tran string `json:"tran"`
				} `json:"words"`
			} `json:"rel"`
		} `json:"rels"`
	} `json:"rel_word"`
	AuthSentsPart struct {
		SentenceCount int    `json:"sentence-count"`
		More          string `json:"more"`
		Sent          []struct {
			Score      float64 `json:"score"`
			Speech     string  `json:"speech"`
			SpeechSize string  `json:"speech-size"`
			Source     string  `json:"source"`
			URL        string  `json:"url"`
			Foreign    string  `json:"foreign"`
		} `json:"sent"`
	} `json:"auth_sents_part"`
	MediaSentsPart struct {
		SentenceCount int    `json:"sentence-count"`
		More          string `json:"more"`
		Query         string `json:"query"`
		Sent          []struct {
			Mediatype string `json:"@mediatype"`
			Snippets  struct {
				Snippet []struct {
					StreamURL string `json:"streamUrl"`
					Duration  string `json:"duration"`
					Swf       string `json:"swf"`
					Name      string `json:"name"`
					Source    string `json:"source"`
					Win8      string `json:"win8"`
				} `json:"snippet"`
			} `json:"snippets"`
			SpeechSize string `json:"speech-size,omitempty"`
			Eng        string `json:"eng"`
			Chn        string `json:"chn,omitempty"`
		} `json:"sent"`
	} `json:"media_sents_part"`
	Etym struct {
		Etyms struct {
			Zh []struct {
				Source string `json:"source"`
				Word   string `json:"word"`
				Value  string `json:"value"`
				URL    string `json:"url"`
				Desc   string `json:"desc"`
			} `json:"zh"`
		} `json:"etyms"`
		Word string `json:"word"`
	} `json:"etym"`
	Special struct {
		Summary struct {
			Sources struct {
				Source struct {
					Site string `json:"site"`
					URL  string `json:"url"`
				} `json:"source"`
			} `json:"sources"`
			Text string `json:"text"`
		} `json:"summary"`
		CoAdd   string `json:"co-add"`
		Total   string `json:"total"`
		Entries []struct {
			Entry struct {
				Major string `json:"major"`
				Trs   []struct {
					Tr struct {
						Nat      string `json:"nat"`
						ChnSent  string `json:"chnSent"`
						Cite     string `json:"cite"`
						DocTitle string `json:"docTitle"`
						EngSent  string `json:"engSent"`
						URL      string `json:"url"`
					} `json:"tr,omitempty"`
				} `json:"trs"`
				Num int `json:"num"`
			} `json:"entry"`
		} `json:"entries"`
	} `json:"special"`
	Senior struct {
		EncryptedData string `json:"encryptedData"`
		Source        struct {
			Name string `json:"name"`
		} `json:"source"`
	} `json:"senior"`
	Input string `json:"input"`
	Meta  struct {
		Input           string   `json:"input"`
		GuessLanguage   string   `json:"guessLanguage"`
		IsHasSimpleDict string   `json:"isHasSimpleDict"`
		Le              string   `json:"le"`
		Lang            string   `json:"lang"`
		Dicts           []string `json:"dicts"`
	} `json:"meta"`
	Le            string `json:"le"`
	OxfordAdvance struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxfordAdvance"`
}
