### Упражнение 7.17

Расширьте возможности ```xmlselect``` так, чтобы элементы могли быть выбраны не только 
по имени, но и по аттрибутам, в духе CSS, так что, например, элемент наподобие
"<div id="page" class="wide">" может быть выбран как по соответствию аттрибутов id 
или class, так и по его имени. 

Решение:
```shell
$ ./fetch http://www.w3.org/TR/2006/RECxml1120060816 | ./xmlselect2 id class
Element: id | Value:www-w3-org
Element: class | Value:w3c_public
Element: id | Value:w3c_container
Element: id | Value:w3c_mast
Element: class | Value:logo
Element: class | Value:alt-logo
Element: id | Value:w3c_nav
Element: class | Value:w3c_sec_nav
Element: class | Value:main_nav
Element: class | Value:first-item
Element: class | Value:last-item
Element: class | Value:search-item
Element: id | Value:search-form
Element: class | Value:text
Element: id | Value:search-submit
Element: class | Value:submit
Element: id | Value:w3c_main
Element: id | Value:w3c_logo_shadow
Element: class | Value:w3c_leftCol
Element: id | Value:w3c_sidenavigation
Element: class | Value:w3c_leftCol
Element: class | Value:offscreen
Element: class | Value:category
Element: class | Value:ribbon
Element: class | Value:header-link
Element: class | Value:theme
Element: class | Value:w3c_mainCol
Element: id | Value:w3c_crumbs
Element: id | Value:w3c_crumbs_frame
Element: class | Value:bct
Element: class | Value:skip
Element: class | Value:cr
Element: class | Value:cr
Element: class | Value:cr
Element: class | Value:current
Element: class | Value:title
Element: id | Value:w3c_content_body
Element: class | Value:line
Element: class | Value:intro tPadding
Element: class | Value:show_items
Element: id | Value:w3c_footer
Element: id | Value:w3c_footer-inner
Element: class | Value:offscreen
Element: class | Value:w3c_footer-nav
Element: class | Value:footer_top_nav
Element: class | Value:last-item
Element: class | Value:w3c_footer-nav
Element: class | Value:footer_bottom_nav
Element: id | Value:w3c_signature
Element: class | Value:w3c_footer-nav
Element: class | Value:footer_follow_nav
Element: class | Value:social-icon
Element: class | Value:copyright

```