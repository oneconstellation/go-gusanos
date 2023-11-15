function byipm.init()
	-- Connect Menu
	gui_load_gss("menu-common")
	gui_load_gss("connect-menu")
	gui_load_gss("byip-menu")
	
	--[[
	local menu = gui_load_xml("byip-menu")
	local win = menu:child("byip-win")
	local connectBtn = win:child("connect")
	local backBtn = win:child("back")
	local ipfield = win:child("ipfield")
	]]
	
	local menu = gui_group({id = "menu", group = "byip"})
	local win = gui_window({id = "win"})
	local connectBtn = gui_button({id = "connect", label="Connect"})
	local backBtn = gui_button({id = "back", label="Back"})
	local ipfield = gui_edit({id = "ipfield"})
	win:add(connectBtn, backBtn, ipfield)
	menu:add(win)
	gui_root():add(menu)
		
	--[[
	<group id="menu" group="byip">
	<window id="win">
		<button id="connect" label="Connect" />
		<button id="back" label="Back" />
		<edit id="ipfield" />
	</window>
</group>
]]

	function byipm.isShown()
		return menu:is_visible()
	end

	function byipm.show()
		if not menu:is_visible() then
			menu:set_visibility(true)
			menu:focus()
		end
	end

	function byipm.hide()
		menu:set_visibility(false)
	end

	function connectBtn:onAction()
		local ip = ipfield:text()
		if ip then
			connect(ip)
			byipm.hide()
			mainm.show()
		end
	end
	
	function backBtn:onAction()
		byipm.hide()
		connectm.show()
	end

	function ipfield:onAction()
		local ip = ipfield:text()
		if ip then
			connect(ip)
		end
	end
	

	function menu:onKeyDown(k)
		if k == Keys.ESC and map_is_loaded() then
			byipm.hide()
			connectm.show()
			return true
		end
	end

	byipm.hide()

end