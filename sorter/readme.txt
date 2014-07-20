P65，排序算法比较程序，从命令行指定输入的数据文件和输出文件，并指定对应的排序算法。
该程序的用法如下所示：
USAGE：sorter -i <in> -o <out> -a <qsort|bubblesort>
例如：
$ ./sorter -i in.dat -o out.dat -a qsort
The sorting process costs 10us to complete.

输入不合法时应该给出对应的提示。