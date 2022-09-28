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

#include <linux/sched.h>

#include <linux/sched/signal.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo para medir uso del CPU");
MODULE_AUTHOR("Fernando Mauricio Gomez Santos");


struct task_struct * cpu;
struct task_struct * child;
struct list_head * lstProcess;

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int escribir_archivo(struct seq_file *archivo, void *v)
{   
     seq_printf(archivo, "[");
     seq_printf(archivo, "\n");
     for_each_process(cpu){
        uid_t uid = __kuid_val(task_uid(cpu));
                seq_printf(archivo, "{\n\"pid\": \"%d\",\n", cpu->pid);
        seq_printf(archivo, "\"comm\": \" %s\",\n", cpu->comm);
        seq_printf(archivo, "\"state\": \"%u\",\n", cpu->__state);
        seq_printf(archivo, "\"owner\": \"%d\",\n", uid);
        seq_printf(archivo, "\"child\": [");
        seq_printf(archivo, "\n");
        list_for_each(lstProcess, &(cpu->children)){
            child = list_entry(lstProcess, struct task_struct, sibling);
            uid_t uid_c = __kuid_val(task_uid(child));
            seq_printf(archivo, "   ");
            seq_printf(archivo, "{\"pid\": \"%d\",\n", child->pid);
            seq_printf(archivo, "   ");
            seq_printf(archivo, "\"comm\": \"%s\",\n", child->comm);
            seq_printf(archivo, "   ");
            seq_printf(archivo, "\"state\": \"%u\",\n", cpu->__state);
            seq_printf(archivo, "   ");
            seq_printf(archivo, "\"owner\": \"%d\"\n", uid_c);
            seq_printf(archivo, "   ");
            seq_printf(archivo, "},\n");
        }
    seq_printf(archivo, "]");
    seq_printf(archivo, "},\n");
    }
    seq_printf(archivo, "]");
    seq_printf(archivo, "\n");
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
    proc_create("cpu201901849", 0, NULL, &operaciones);
    printk(KERN_INFO "*********************************************\n");
    printk(KERN_INFO "*********************************************\n");
    printk(KERN_INFO "***************Fernando Gomez****************\n");
    printk(KERN_INFO "*********************************************\n");
    printk(KERN_INFO "*********************************************\n");
    return 0;
}

//Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("cpu201901849", NULL);
    printk(KERN_INFO "*********************************************\n");
    printk(KERN_INFO "*********************************************\n");
    printk(KERN_INFO "****************Segundo Semestre*************\n");
    printk(KERN_INFO "*********************************************\n");
    printk(KERN_INFO "*********************************************\n");
}

module_init(_insert);
module_exit(_remove);