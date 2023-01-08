import pandas as pd 
df = pd.read_json("/Users/skm/Documents/Code/Go/nascar/nascar2013_2021.json")
# print(df)
print(df['Season'].value_counts())
df2 = df[df['Finish']!=""]
print(df['Season'].value_counts(), df2['Season'].value_counts())