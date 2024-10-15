package mysql

import "fmt"

func DefaultSettings() {
	var d = "Discover the charm of our handcrafted clothing, where every piece tells a story. Made with love and attention to detail, our unique garments celebrate individuality and style. Embrace the beauty of sustainable fashion and elevate your wardrobe with designs that stand out. Shop now and wear your story!"

	var query = fmt.Sprintf("INSERT INTO settings VALUES('hero_one_description','%s'),('hero_two_description','%s'),('main_description','%s'),('hero_one_title', '%s'),('hero_two_title','%s');", d, d, d, "title1", "title2")

	Db.Exec(query)
}

func UpdateMainDescription(description string) error {
	var query = fmt.Sprintf("UPDATE settings SET value = '%s' WHERE setting = 'main_description';", description)

	var _, err = Db.Exec(query)

	return err
}

func UpdateHeroOne(title string, description string) error {
	var query1 = fmt.Sprintf("UPDATE settings SET value = '%s' WHERE setting = 'hero_one_description';", description)
	var query2 = fmt.Sprintf("UPDATE settings SET value = '%s' WHERE setting = 'hero_one_title';", title)

	var _, err = Db.Exec(query1)
	_, err = Db.Exec(query2)

	return err
}

func UpdateHeroTwo(title string, description string) error {
	var query1 = fmt.Sprintf("UPDATE settings SET value = '%s' WHERE setting = 'hero_two_description';", description)
	var query2 = fmt.Sprintf("UPDATE settings SET value = '%s' WHERE setting = 'hero_two_title';", title)

	var _, err = Db.Exec(query1)
	_, err = Db.Exec(query2)

	return err
}

func GetSettings() map[string]string {
	var settings = map[string]string{}
	var r, _ = Db.Query("SELECT setting,value FROM settings")
	for r.Next() {
		var k string
		var v string
		r.Scan(&k, &v)

		settings[k] = v
	}

	return settings
}
