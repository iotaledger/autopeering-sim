import numpy as np
import matplotlib.pyplot as plt
import argparse

xcol = 0
xlabel = "Time [seconds]"
xscale = 'linear'

# - - - - parameters unlikely to be changed - - -
ycol = 1 
zcol = 2  # last column in the csv file
folder = "data/"


def main():
    # argument parser settings
    parser = argparse.ArgumentParser()
    parser.add_argument('--N', type=int, required=True, help='minimum number of nodes')
    parser.add_argument('--maxN', type=int, required=True,help='maximum number of nodes')
    parser.add_argument('--t', type=int, required=True, help='interval of simulation')
    args = parser.parse_args()
    interval = (args.maxN - args.N) // args.t + 1 
    # Plot 
    #printLinkAnalysis(args, interval)
    #printConvergenceAnalysis()
    printDistanceAnalysis(args, interval)


def printLinkAnalysis(args, interval):
    fig = plt.figure()
    filename = folder+'plot_linkAnalysis'
    for i in range(interval):
        csv = "linkAnalysis_" + str(i)
        label = "N=" + str(args.N + i * args.t) 
        cdfPlot2("LinkAnalysis", csv, filename, label)
    plt.xscale(xscale)
    #plt.xlim(xlim)
    plt.xlabel(xlabel)
    plt.ylabel("Probability")
    #plt.yscale('log')
    plt.legend(loc='best')
    plt.savefig(filename+'.eps', format='eps')
    plt.clf()

def printConvergenceAnalysis(): 
    filename = folder+'plot_convAnalysis'
    partPlot3("ConvAnalysis", "convAnalysis", filename, "blue")

def printDistanceAnalysis(args, interval):
    fig = plt.figure()
    filename = folder+'plot_distanceAnalysis'
    ary = [33, 100, 330, 1000]
    for i in range(len(ary)):
        csv = "distanceAnalysis_" + str(i)
        #label = "N=" + str(args.N + i * args.t) 
        label = "N=" + str(ary[i]) 
        logPlot2("DistanceAnalysis", csv, filename, label)
        #cdfPlot2("DistanceAnalysis", csv, filename, label)
    plt.xlabel("Distance")
    plt.ylabel("Probability log(1 - CDF)")
    #plt.ylabel("Probability (CDF)")
    plt.legend(loc='best')
    plt.savefig(filename+'.eps', format='eps')
    plt.clf()

def cdfPlot2(type, file, filename, label):
    x = loadDatafromRow(file, xcol)
    y = loadDatafromRow(file, ycol)
    x, y = sort2vecs(x, y)
    CY = np.cumsum(y)
    plt.plot(x, CY, linewidth=1, label=label)
    np.savez(filename+"_"+type, x=x, y=y)

def logPlot2(type, file, filename, label):
    x = loadDatafromRow(file, xcol)
    y = loadDatafromRow(file, ycol)
    x, y = sort2vecs(x, y)
    CY = np.cumsum(y)
    #CY = CY[CY <= 0.3]
    #x = x[:len(CY)]
    CY = 1.0-CY
    plt.yscale('log')
    # ignore the last element log(0)
    plt.plot(x[:-1], CY[:-1], linewidth=1, label=label)
    np.savez(filename+"_"+type, x=x, y=y)

def partPlot2(type, file, filename, label):
    x = loadDatafromRow(file, xcol)
    y = loadDatafromRow(file, ycol)
    x, y = sort2vecs(x, y)
    plt.plot(x, y, linewidth=1, label=label)
    np.savez(filename+"_"+type, x=x, y=y)

def partPlot3(type, file, filename, color):
    x = loadDatafromRow(file, xcol)
    y = loadDatafromRow(file, ycol)
    z = loadDatafromRow(file, zcol)
    x, y, z = sort3vecs(x, y, z)
    
    fig, ax1 = plt.subplots()
    
    color = 'tab:blue'
    ax1.set_xlabel('Time [seconds]')
    ax1.set_ylabel('Nodes with 8 neighbors [%]', color=color)
    ax1.set_ylim([0, 100])
    ax1.plot(x, y, color=color)
    ax1.tick_params(axis='y', labelcolor=color)

    ax2 = ax1.twinx()  # instantiate a second axes that shares the same x-axis

    color = 'tab:red'
    ax2.set_ylabel('Avg # of neighbors', color=color)  # we already handled the x-label with ax1
    ax2.plot(x, z, color=color)
    ax2.tick_params(axis='y', labelcolor=color)

    fig.tight_layout()  # otherwise the right y-label is slightly clipped

    np.savez(filename+"_"+type, x=x, y=y)
    fig.savefig(filename+'.eps', format='eps')
    fig.clf()

def sort2vecs(x, y):
    i = np.argsort(x)
    x = x[i]
    y = y[i]
    return x, y

def sort3vecs(x, y, z):
    i = np.argsort(x)
    x = x[i]
    y = y[i]
    z = z[i]
    return x, y, z

def loadDatafromRow(datatype, row):
    try:
        filestr = folder+'result_'+datatype+'.csv'
        f = open(filestr, "r")
        data = np.loadtxt(f, delimiter=",", skiprows=1, usecols=(row))
        return data
    except IOError:
        print(filestr)
        print("File not found.")
        return []


# needs to be at the very end of the file
if __name__ == '__main__':
    main()
