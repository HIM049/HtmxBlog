package model

type I18n struct {
	Common struct {
		SwitchTheme   string `json:"switch_theme"`
		Edit          string `json:"edit"`
		Delete        string `json:"delete"`
		Add           string `json:"add"`
		Create        string `json:"create"`
		Upload        string `json:"upload"`
		Remove        string `json:"remove"`
		None          string `json:"none"`
		Public        string `json:"public"`
		Hide          string `json:"hide"`
		Private       string `json:"private"`
		ConfirmDelete string `json:"confirm_delete"`
	} `json:"common"`

	Auth struct {
		Title               string `json:"title"`
		Subtitle            string `json:"subtitle"`
		PasswordLabel       string `json:"password_label"`
		PasswordPlaceholder string `json:"password_placeholder"`
		SubmitButton        string `json:"submit_button"`
		BackToHome          string `json:"back_to_home"`
		Copyright           string `json:"copyright"`
	} `json:"auth"`

	Dashboard struct {
		Title                 string `json:"title"`
		ManagePostsTitle      string `json:"manage_posts_title"`
		ManagePostsDesc       string `json:"manage_posts_desc"`
		ManagePagesTitle      string `json:"manage_pages_title"`
		ManagePagesDesc       string `json:"manage_pages_desc"`
		ManageCategoriesTitle string `json:"manage_categories_title"`
		ManageCategoriesDesc  string `json:"manage_categories_desc"`
		ManageSettingsTitle   string `json:"manage_settings_title"`
		ManageSettingsDesc    string `json:"manage_settings_desc"`
		ManageRedirectsTitle  string `json:"manage_redirects_title"`
		ManageRedirectsDesc   string `json:"manage_redirects_desc"`
		ManageCommentsTitle   string `json:"manage_comments_title"`
		ManageCommentsDesc    string `json:"manage_comments_desc"`
		VisitStatisticsTitle  string `json:"visit_statistics_title"`
		VisitStatisticsDesc   string `json:"visit_statistics_desc"`
	} `json:"dashboard"`

	Posts struct {
		Title         string `json:"title"`
		CreateNewPost string `json:"create_new_post"`
		BackToPosts   string `json:"back_to_posts"`
		ThID          string `json:"th_id"`
		ThTitle       string `json:"th_title"`
		ThCategory    string `json:"th_category"`
		ThState       string `json:"th_state"`
		ThActions     string `json:"th_actions"`
		ConfirmDelete string `json:"confirm_delete"`
	} `json:"posts"`

	Editor struct {
		TitlePlaceholder          string `json:"title_placeholder"`
		ContentPlaceholder        string `json:"content_placeholder"`
		DraftBadge                string `json:"draft_badge"`
		UnpublishedDraftBadge     string `json:"unpublished_draft_badge"`
		SavingIndicator           string `json:"saving_indicator"`
		SaveDraft                 string `json:"save_draft"`
		Publish                   string `json:"publish"`
		ConfirmPublish            string `json:"confirm_publish"`
		MetaInfo                  string `json:"meta_info"`
		Visibility                string `json:"visibility"`
		Protect                   string `json:"protect"`
		ProtectNone               string `json:"protect_none"`
		ProtectLogin              string `json:"protect_login"`
		ProtectPassword           string `json:"protect_password"`
		Category                  string `json:"category"`
		Tags                      string `json:"tags"`
		TagsPlaceholder           string `json:"tags_placeholder"`
		CreatedAt                 string `json:"created_at"`
		CustomVars                string `json:"custom_vars"`
		AddCustomVar              string `json:"add_custom_var"`
		CustomVarKeyPlaceholder   string `json:"custom_var_key_placeholder"`
		CustomVarValuePlaceholder string `json:"custom_var_value_placeholder"`
	} `json:"editor"`

	Attachments struct {
		Title         string `json:"title"`
		SelectFile    string `json:"select_file"`
		Upload        string `json:"upload"`
		CopyLink      string `json:"copy_link"`
		Copied        string `json:"copied"`
		Remove        string `json:"remove"`
		ConfirmRemove string `json:"confirm_remove"`
	} `json:"attachments"`

	Categories struct {
		Title           string `json:"title"`
		CreateTitle     string `json:"create_title"`
		NameLabel       string `json:"name_label"`
		NamePlaceholder string `json:"name_placeholder"`
		ColorLabel      string `json:"color_label"`
		PickColorHint   string `json:"pick_color_hint"`
		CreateButton    string `json:"create_button"`
		ExistingTitle   string `json:"existing_title"`
		NoCategories    string `json:"no_categories"`
		HiddenTag       string `json:"hidden_tag"`
		PrivateTag      string `json:"private_tag"`
		ConfirmDelete   string `json:"confirm_delete"`
	} `json:"categories"`

	Pages struct {
		Title               string `json:"title"`
		SectionTitle        string `json:"section_title"`
		CreateTitle         string `json:"create_title"`
		NameLabel           string `json:"name_label"`
		NamePlaceholder     string `json:"name_placeholder"`
		RouteLabel          string `json:"route_label"`
		RoutePlaceholder    string `json:"route_placeholder"`
		TemplateLabel       string `json:"template_label"`
		TemplatePlaceholder string `json:"template_placeholder"`
		CreateButton        string `json:"create_button"`
		ConfirmDelete       string `json:"confirm_delete"`
	} `json:"pages"`

	Redirects struct {
		Title                 string `json:"title"`
		CreateTitle           string `json:"create_title"`
		SourcePathLabel       string `json:"source_path_label"`
		SourcePathPlaceholder string `json:"source_path_placeholder"`
		TargetPathLabel       string `json:"target_path_label"`
		TargetPathPlaceholder string `json:"target_path_placeholder"`
		CodeLabel             string `json:"code_label"`
		AddButton             string `json:"add_button"`
		ExistingTitle         string `json:"existing_title"`
		NoRedirects           string `json:"no_redirects"`
		SourceCol             string `json:"source_col"`
		TargetCol             string `json:"target_col"`
		ConfirmDelete         string `json:"confirm_delete"`
	} `json:"redirects"`

	Settings struct {
		Title            string `json:"title"`
		CreateTitle      string `json:"create_title"`
		KeyLabel         string `json:"key_label"`
		KeyPlaceholder   string `json:"key_placeholder"`
		ValueLabel       string `json:"value_label"`
		ValuePlaceholder string `json:"value_placeholder"`
		AddButton        string `json:"add_button"`
		ExistingTitle    string `json:"existing_title"`
		NoSettings       string `json:"no_settings"`
		ConfirmDelete    string `json:"confirm_delete"`
	} `json:"settings"`

	Comments struct {
		Title         string `json:"title"`
		ExistingTitle string `json:"existing_title"`
		ThAuthor      string `json:"th_author"`
		ThContent     string `json:"th_content"`
		ThDate        string `json:"th_date"`
		ThActions     string `json:"th_actions"`
		NoComments    string `json:"no_comments"`
		Approve       string `json:"approve"`
		Delete        string `json:"delete"`
		ConfirmDelete string `json:"confirm_delete"`
	} `json:"comments"`

	Statistics struct {
		Title             string `json:"title"`
		TotalPVTitle      string `json:"total_pv_title"`
		TotalPVDesc       string `json:"total_pv_desc"`
		TotalUVTitle      string `json:"total_uv_title"`
		TotalUVDesc       string `json:"total_uv_desc"`
		ArticleStatsTitle string `json:"article_stats_title"`
		NoArticleData     string `json:"no_article_data"`
		SourceStatsTitle  string `json:"source_stats_title"`
		NoSourceData      string `json:"no_source_data"`
	} `json:"statistics"`
}
