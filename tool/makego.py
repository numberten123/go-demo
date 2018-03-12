# -*- coding: utf-8 -*-
import xlrd
import sys

reload(sys)   
sys.setdefaultencoding('utf8')  

EXPORT_TYPE_ROW  = 2
EXPORT_VALUE_NAME_ROW = 3
EXPORT_VALUE_TYPE_ROW = 4

def excle_to_go(excle_file,export_type,output_dir='./output/'):
	data = xlrd.open_workbook(excle_file)
	for sheet in data.sheets():
		table_name = sheet.cell(0,0).value
		Table_name = table_name.capitalize()
		all=[]
		erl_table=[]
		temp1_str=''
		count=0
		for r in range(5,sheet.nrows):
			line_value =[]
			count=0
			for c in range(0,sheet.ncols):
				if sheet.cell(EXPORT_TYPE_ROW,c).value.find(export_type) > -1 :
					count=count+1
					if c == 0:
						v = sheet.cell(r,c).value
						if v==None or v == '':
							v=0
						if is_number(v) ==True:
							v = int(v)
						all.append(str(v))
						temp1_str=str(v)
						line_str = '\n\t\t'+temp1_str+':\n\t\t\tgo_include.'+Table_name+'{\n\t\t\t\t'
					feildname = sheet.cell(EXPORT_VALUE_NAME_ROW,c).value
					value_str = ''+ feildname.capitalize()+ ':'
					_type = sheet.cell(EXPORT_VALUE_TYPE_ROW,c).value
					if _type  == 'string':
						str1 = str(sheet.cell(r,c).value)
						if is_number(str1) ==True:
							str1 = int(float(str1))
						nowstr = '"' + str(str1) + '"'
						value_str += nowstr.replace("\n", "\\n")+ ','
					elif _type == 'number':
						v = sheet.cell(r,c).value
						if v==None or v == '':
							v=0
						if (v == round(v)):
							v = int(v)
						value_str += str(v)+ ','
					else:
						nowstr = '' + str(sheet.cell(r,c).value) + ''
						value_str += nowstr+ ','
					line_value.append(value_str)
			v = '\n\t\t\t\t'.join(line_value)
			line_str = line_str+v+'\n\t\t\t},'
			erl_table.append(line_str)
		
		str1_str=','.join(all)
		erl =  '// !!! DO NOT EDIT !!!\n// auto generated from Excel files\npackage go_config\n\nimport(\n\t"test/go_include"\n\t"test/go_log"\n)\n\nfunc '+Table_name+'(key uint32) (go_include.'+Table_name+') {\n\tvar '+table_name+' = map[uint32]go_include.'+Table_name+'{'+''.join(erl_table)+'\n\t}\n\tif result, ok := '+table_name+'[key]; ok {\n\t\treturn result\n\t}\n\tgo_log.Error("'+Table_name+' not match:", key)\n\treturn go_include.'+Table_name+'{}\n}'
		f = open(output_dir + table_name + '.go','wb')
		f.write(erl)
		f.close()
	
	return 


def is_number(s):
    try:
        float(s)
        return True
    except ValueError:
        pass
    return False

# excle_to_hrl_define('error_code.xlsx')
try:  
    excle_to_go(sys.argv[1],'s') 
except Exception,e:
    print "ignore :",sys.argv[1]

