import numpy as np
import matplotlib.pyplot as plt
import csv
import pandas as pd

xcol = 0
xlabel = 'Time [seconds]'
xscale = 'linear'

# - - - - parameters unlikely to be changed - - -
ycol = 1
zcol = 2  # last column in the csv file
folder = 'data/'


def main():
    # printLinkAnalysis()
    printConvergenceAnalysis()
    printCountMsgAnalysis('In')
    printCountMsgAnalysis('Out')
    # printMsgAnalysis()
    # printMsgPerTAnalysis()


def printLinkAnalysis():
    fig = plt.figure()
    filename = folder+'plot_linkAnalysis'
    partPlot2('LinkAnalysis', 'linkAnalysis', filename, 'blue')
    plt.xscale(xscale)
    # plt.xlim(xlim)
    plt.xlabel(xlabel)
    plt.ylabel('Probability')
    # plt.yscale('log')
    plt.legend(loc='best')
    plt.savefig(filename+'.eps', format='eps')
    plt.clf()


def printConvergenceAnalysis():
    filename = folder+'plot_convAnalysis'
    partPlot3('ConvAnalysis', 'convAnalysis', filename, 'blue')


def printCountMsgAnalysis(type):
    fig = plt.figure()
    filename = folder+'plot_convMsg'+type+'Analysis'
    histoPlot1('CountMsg'+type+'Analysis', 'countMsg' +
               type+'Analysis', filename, 'blue', xcol)
    plt.xscale(xscale)
    plt.xlabel('# of messages')
    plt.ylabel('# of nodes')
    plt.title(type+'bound')
    plt.savefig(filename+'.eps', format='eps')
    plt.clf()


def printMsgAnalysis():
    fig = plt.figure()
    filename = folder+'plot_MsgAnalysis'
    histoPlot1('vMsgAnalysis', 'MsgAnalysis', filename, 'blue', 1)
    plt.xscale(xscale)
    plt.xlabel('# of messages')
    plt.ylabel('# of nodes')
    plt.savefig(filename+'.eps', format='eps')
    plt.clf()


def printMsgPerTAnalysis():
    fig = plt.figure()
    filename = folder+'plot_msgPerTAnalysis'
    histoPlotMulti('msgPerTAnalysis', 'msgPerTAnalysis', filename, 'blue')
    plt.xscale(xscale)
    plt.xlabel('# of messages per T (dropAll=true)')
    plt.ylabel('# of nodes')
    plt.savefig(filename+'.eps', format='eps')
    plt.clf()


def histoPlotMulti(type, file, filename, color):
    color = 'tab:blue'
    bandwidth = 1
    x = loadDatafromRow(file, 2)
    for i in range(3, 10):
        t = loadDatafromRow(file, i)
        x = np.append(x, t)
    plt.hist(x, bins=range(int(np.amin(x)), int(np.amax(x)), bandwidth))
    axes = plt.gca()
    axes.set_xlim([0, 100])
    np.savez(filename+'_'+type, x=x)


def histoPlot1(type, file, filename, color, valcol):
    color = 'tab:blue'
    x = loadDatafromRow(file, valcol)

    # save the bins
    y = binX(x)
    pd.DataFrame(y).to_csv('data/plot_convMsgAnalysis_histdata.csv')
    # writer = csv.writer(open('test', 'w'))
    # for row in x:
    #     writer.writerow(row)

    xmax = np.amax([int(np.amax(x)), 100])
    maxx = np.amax([np.amin(x)+1, np.amax(x)])
    bandwidth = 1
    plt.hist(x, bins=range(int(np.amin(x))-1, int(maxx), bandwidth))
    axes = plt.gca()
    axes.set_xlim([0, xmax])
    np.savez(filename+'_'+type, x=x)


def partPlot2(type, file, filename, color):
    color = 'tab:blue'
    x = loadDatafromRow(file, xcol)
    y = loadDatafromRow(file, ycol)
    x, y = sort2vecs(x, y)
    plt.plot(x, y, color=color, linewidth=1)
    np.savez(filename+'_'+type, x=x, y=y)


def partPlot3(type, file, filename, color):
    x = loadDatafromRow(file, xcol)
    y = loadDatafromRow(file, ycol)
    z = loadDatafromRow(file, zcol)
    x, y, z = sort3vecs(x, y, z)

    fig, ax1 = plt.subplots()

    color = 'tab:blue'
    ax1.set_xlabel('Time [seconds]')
    ax1.set_ylabel('Nodes with 8 neighbors [%]', color=color)
    ax1.set_ylim([50, 100])
    ax1.set_xlim([0, max(x)])
    ax1.plot(x, y, color=color)
    ax1.spines['top'].set_visible(False)
    ax1.spines['left'].set_color('grey')
    ax1.spines['right'].set_color('grey')
    ax1.tick_params(axis='y', labelcolor=color)
    # Show the grid lines as dark grey lines
    plt.grid(b=True, which='major', color='#666666', linestyle='-')

    ax2 = ax1.twinx()  # instantiate a second axes that shares the same x-axis

    color = 'tab:red'
    # we already handled the x-label with ax1
    ax2.set_ylabel('Avg # of neighbors', color=color)
    ax2.plot(x, z, color=color)
    ax2.set_ylim([6, 8])
    ax2.set_xlim([0, max(x)])
    ax2.spines['top'].set_color('grey')
    ax2.spines['left'].set_color('grey')
    ax2.spines['right'].set_color('grey')
    ax2.tick_params(axis='y', labelcolor=color)

    fig.tight_layout()  # otherwise the right y-label is slightly clipped

    np.savez(filename+'_'+type, x=x, y=y)
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
        f = open(filestr, 'r')
        data = np.loadtxt(f, delimiter=',', skiprows=1, usecols=(row))
        return data
    except IOError:
        print(filestr)
        print('File not found.')
        return []


def binX(X):
    bins = np.zeros(int(np.amax(X))+1)
    for x in X:
        bins[int(x)] += 1
    return bins


# needs to be at the very end of the file
if __name__ == '__main__':
    main()
