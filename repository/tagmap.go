// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

type TagMap map[Tag]ItemList

func NewTagMap() TagMap {
	return make(TagMap)
}

func (tagmap TagMap) Add(item *Item) {

	for _, tag := range item.MetaData.Tags {

		if itemlist, exists := tagmap[tag]; exists {

			// add the item to the item list for this tag
			tagmap[tag] = itemlist.Add(item)

		} else {

			// create a new item list
			tagmap[tag] = NewItemList(item)
		}

	}

}

func (tagmap TagMap) Remove(item *Item) {

	for tag, itemlist := range tagmap {

		tagmap[tag] = itemlist.Remove(item)

		// remove tag if item list is empty
		if itemlist.IsEmpty() {
			delete(tagmap, tag)
		}
	}

}
