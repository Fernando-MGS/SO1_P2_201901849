#include <linux/module.h>
#define INCLUDE_VERMAGIC
#include <linux/build-salt.h>
#include <linux/elfnote-lto.h>
#include <linux/vermagic.h>
#include <linux/compiler.h>

BUILD_SALT;
BUILD_LTO_INFO;

MODULE_INFO(vermagic, VERMAGIC_STRING);
MODULE_INFO(name, KBUILD_MODNAME);

__visible struct module __this_module
__section(".gnu.linkonce.this_module") = {
	.name = KBUILD_MODNAME,
	.init = init_module,
#ifdef CONFIG_MODULE_UNLOAD
	.exit = cleanup_module,
#endif
	.arch = MODULE_ARCH_INIT,
};

#ifdef CONFIG_RETPOLINE
MODULE_INFO(retpoline, "Y");
#endif

static const struct modversion_info ____versions[]
__used __section("__versions") = {
	{ 0xcb440b5e, "module_layout" },
	{ 0x6b6889b6, "seq_read" },
	{ 0xa93bfc46, "remove_proc_entry" },
	{ 0x92997ed8, "_printk" },
	{ 0xe832bcc7, "proc_create" },
	{ 0x2d5f69b3, "rcu_read_unlock_strict" },
	{ 0xe4133be1, "init_task" },
	{ 0xe7a4bdfe, "seq_printf" },
	{ 0x5b8239ca, "__x86_return_thunk" },
	{ 0xc1025998, "single_open" },
	{ 0xbdfb6dbb, "__fentry__" },
};

MODULE_INFO(depends, "");


MODULE_INFO(srcversion, "4CECFFAE5F77B339030B6AE");
