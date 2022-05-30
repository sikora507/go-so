using BenchmarkDotNet.Attributes;

public class Benchmarks
{
    [Benchmark]
    public string TestIntToString(){
        string s = "";
        for (int i = 0; i < 100000; i++)
        {
            s = StringUtils.IntToString(i);
        }
        return s;
    }
}