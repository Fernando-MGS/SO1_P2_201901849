#include <linux/module.h>
// para usar KERN_INFO
#include <linux/kernel.h>

//Header para los macros module_init y module_exit
#include <linux/init.h>
//Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
/* for copy_from_user */
#include <asm/uaccess.h>        
/* Header para usar la lib seq_file y manejar el archivo en /proc*/
#include <linux/seq_file.h>

#include <linux/sys.h>

#include <linux/hugetlb.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo para medir uso de la RAM");
MODULE_AUTHOR("Fernando Mauricio Gomez Santos");

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int escribir_archivo(struct seq_file *archivo, void *v)
{   
    struct sysinfo si;
    si_meminfo(&si);
    seq_printf(archivo, "{\n");
    seq_printf(archivo, "\"free\": %lu\n",si.freeram);
    seq_printf(archivo, "\"total\": %lu\n",si.totalram);
    seq_printf(archivo, "}");
    return 0;
}

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

//Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int _insert(void)
{
    proc_create("ram_201901849", 0, NULL, &operaciones);
    printk(KERN_INFO "*******************************\n");
    printk(KERN_INFO "*******************************\n");
    printk(KERN_INFO "************201901849**********\n");
    printk(KERN_INFO "*******************************\n");
    printk(KERN_INFO "*******************************\n");
    return 0;
}

//Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("ram_201901849", NULL);
    printk(KERN_INFO "*******************************\n");
    printk(KERN_INFO "*******************************\n");
    printk(KERN_INFO "*****SISTEMAS-OPERATIVOS-1*****\n");
    printk(KERN_INFO "*******************************\n");
    printk(KERN_INFO "*******************************\n");
}

module_init(_insert);

module_exit(_remove);