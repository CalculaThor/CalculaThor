#include <iostream>
#include <math.h>
#include <vector>
#include <cstdio>
#include "exprtk.hpp"

struct rootRange {
    long double a, b, root;
    bool isRoot, isRange;
    unsigned int iterations;
    rootRange() {
        a = 0;
        b = 0;
        root = 0;
        isRoot =  false;
        isRange = false;
        iterations = 0;
    };
};

struct rootAnswer {
    long double root;
    bool isRoot, isAlmostRoot, badIn;
    unsigned int iterations;
    rootAnswer() {
        root = 0;
        isRoot =  false;
        isAlmostRoot = false;
        badIn = false;
        iterations = 0;
    };
};

struct reg1 {
    long double xi, xs, xm, fxm;
    int it;
};

struct reg2 {
    long double xm, fxm;
    int it;
};

struct reg3 {
    long double xm, fxm, dfxm, d2fxm;
    int it;
};

struct reg4 {
    long double xm, fxm, dfxm;
    int it;
};

std::vector<reg1> bisTable;
std::vector<reg1> falTable;
std::vector<reg2> fixTable;
std::vector<reg2> secTable;
std::vector<reg4> newTable;
std::vector<reg3> mulTable;


long double X;

typedef exprtk::expression<long double> expression;
typedef exprtk::parser<long double> parser;

expression func, gfunc, der, der2;

//std::string F, dF, d2F;

void initParser() {
    typedef exprtk::symbol_table<long double> symbolTable;
    typedef exprtk::parser<long double> parser;

    symbolTable sym;
    sym.add_variable("x",X);
    func.register_symbol_table(sym);
    gfunc.register_symbol_table(sym);
    der.register_symbol_table(sym);
    der2.register_symbol_table(sym);

    parser p1, p2, p3, p4;
    p1.compile(F, func);
    p2.compile(dF, der);
    p3.compile(d2F, der2);
    p4.compile(G, gfunc);
}

bool setF(std::string f) {
    parser p;
    return p.compile(f, func)
}

bool setDF(std::string df) {
    parser p;
     return p.compile(df, der)
}

bool setD2F(std::string d2f) {
    parser p;
    return p.compile(d2f, der2)
}

bool setG(std::string g) {
    parser p;
    return p.compile(g, gfunc)
}
 
long double f(long double x) {
    X = x;
    return func.value();
}

long double df(long double x) {
    X = x;
    return der.value();
}

long double d2f(long double x) {
    X = x;
    return der2.value();
}

long double g(long double x) {
    X = x;
    return gfunc.value();
}



rootRange incrementalSearch(long double x0, long double dx, unsigned int maxIt) {
    rootRange ans;

    long double xi = x0;
    long double yi = f(xi);
    if(yi == 0) {
        ans.root = xi;
        ans.isRoot = true;
        return ans;
    }
    long double xf = xi + dx;
    long double yf = f(xf);
    unsigned int count = 1;

    while(yi * yf > 0 && count <= maxIt) {
        xi = xf;
        yi = yf;
        xf = xi + dx;
        yf = f(xf);
        count++;
    }
    ans.iterations = count;

    if(yf == 0) {
        ans.root = xf;
        ans.isRoot = true;
        return ans;
    }
    if(yi * yf < 0) {
        if(xi > xf) {
            long double temp = xi;
            xi = xf;
            xf = temp;
        }
        ans.a = xi;
        ans.b = xf;
        ans.isRange = true;
        return ans;
    }

    return ans;
}

rootAnswer bisection(long double a, long double b, unsigned int maxIt, long double tolerance) {
    rootAnswer ans;
    long double xi = a;
    long double yi =f(xi);
    if(yi == 0) {
        ans.root = xi;
        ans.isRoot = true;
        return ans;
    }

    long double xs = b;
    long double ys = f(xs);
    if(ys == 0) {
        ans.root = xs;
        ans.isRoot = true;
        return ans;
    }
    if(yi*ys > 0) {
        ans.badIn =  true;
        return ans;
    }
    long double xm = (xi + xs) / 2;
    long double ym = f(xm);
    long double err = INFINITY;
    unsigned int count = 1;
    long double temp;
    bisTable.push_back({xi, xs, xm, ym, (int)count});
    while(ym != 0 && err > tolerance && count < maxIt) {
        if(yi*ym < 0) {
            xs = xm;
            ys = ym;
        } else {
            xi = xm;
            yi = ym;
        }
        temp = xm;
        xm = (xi + xs) / 2;
        ym = f(xm);
        err = abs(xm - temp);
        count++;
        bisTable.push_back({xi, xs, xm, ym, (int)count});
    }
    ans.iterations = count;
    if(ym == 0) {
        ans.isRoot = true;
        ans.root = xm;
        return ans;
    }
    if(err <= tolerance) {
        ans.isAlmostRoot = true;
        ans.root = xm;
        return ans;
    }

    return ans;
}

rootAnswer falsePosition(long double a, long double b, unsigned int maxIt, long double tolerance) {
    rootAnswer ans;
    long double xi = a;
    long double yi =f(xi);
    if(yi == 0) {
        ans.root = xi;
        ans.isRoot = true;
        return ans;
    }

    long double xs = b;
    long double ys = f(xs);
    if(ys == 0) {
        ans.root = xs;
        ans.isRoot = true;
        return ans;
    }
    if(yi*ys > 0) {
        ans.badIn =  true;
        return ans;
    }
    long double xm = xi - (yi * (xs - xi) / (ys - yi));
    long double ym = f(xm);
    long double err = INFINITY;
    unsigned int count = 1;
    long double temp = 0;
    falTable.push_back({xi, xs, xm, ym, (int)count});
    while(ym != 0 && err > tolerance && count < maxIt) {
        if(yi*ym < 0) {
             xs = xm;
             ys = ym;
        } else {
            xi = xm;
            yi = ym;
        }
        temp = xm;
        xm = xi - (yi * (xs - xi) / (ys - yi));
        ym = f(xm);
        err = abs(xm - temp);
        count++;
        falTable.push_back({xi, xs, xm, ym, (int)count});
    }
    ans.iterations = count;
    if(ym == 0) {
        ans.isRoot = true;
        ans.root = xm;
        return ans;
    }
    if(err <= tolerance) {
        ans.isAlmostRoot = true;
        ans.root = xm;
        return ans;
    }

    return ans;
}

rootAnswer fixedPoint(long double x0, unsigned int maxIt, long double tolerance) {
    rootAnswer ans;
    long double xn = x0;
    long double yn = f(xn);
    long double err = INFINITY;
    unsigned int count = 0;
    long double last;
    fixTable.push_back({xn, yn, 0});

    while(yn != 0 && err > tolerance && count < maxIt) {
        last = xn;
        xn = g(xn);
        yn = f(xn);
        err = abs(xn - last);
        count++;
        fixTable.push_back({xn, yn, (int)count});
    }
    ans.iterations = count;
    if(yn == 0) {
        ans.isRoot = true;
        ans.root = xn;
        return ans;
    }
    if(err <= tolerance) {
        ans.isAlmostRoot = true;
        ans.root = xn;
        return ans;
    }
    return ans;
}

rootAnswer newton(long double x0, unsigned int maxIt, long double tolerance) {
    rootAnswer ans;
    long double x = x0;
    long double y = f(x);
    long double dy = df(x);
    unsigned int count = 0;
    long double err = INFINITY;
    long double x1;
    newTable.push_back({x, y, dy, 0});
    while(y != 0 && err > tolerance && dy != 0 && count < maxIt) {
        x1 = x - y/dy;
        y = f(x1);
        dy = df(x1);
        err = abs(x1 - x);
        x = x1;
        count++;
        newTable.push_back({x, y, dy, (int)count});
    }
    ans.iterations = count;
    if(y == 0) {
        ans.isRoot = true;
        ans.root = x;
        return ans;
    }
    if(err <= tolerance) {
        ans.isAlmostRoot = true;
        ans.root = x;
        return ans;
    }
    if(dy == 0) {
        ans.badIn = true;
        return ans;
    }
    return ans;
    
}

rootAnswer secant(long double a, long double b, unsigned int maxIt, long double tolerance) {
    rootAnswer ans;
    long double x0 = a;
    long double y0 = f(x0);
    if(y0 == 0) {
        ans.isRoot = true;
        ans.root = x0;
        return ans;
    }
    long double x1 = b;
    long double y1 = f(x1);
    unsigned int count = 0;
    long double err = INFINITY;
    long double x2;
    while(y1 != 0 && err > tolerance && y1 != y0 && count < maxIt) {
        x2 = x1 - y1 * (x1 - x0) / (y1 - y0);
        err = abs(x2 - x1);
        x0 = x1;
        y0 = y1;
        x1 = x2;
        y1 = f(x1);
        count++;
        secTable.push_back({x1, y1, (int)count});
    }
    ans.iterations = count;
    if(y1 == 0) {
        ans.isRoot = true;
        ans.root = x1;
        return ans;
    }
    if(err <= tolerance) {
        ans.isAlmostRoot =  true;
        ans.root = x1;
        return ans;
    }
    if(y1 == y0) {
        ans.badIn =  true;
    }
    return ans;
}

rootAnswer multipeRoot(long double x0, unsigned int maxIt, long double tolerance) {
    rootAnswer ans;
    long double x = x0;
    long double y = f(x);
    long double dy = df(x);
    long double dy2 = pow(dy, 2);
    long double d2y = d2f(x);
    unsigned int count = 0;
    long double err = INFINITY;
    long double x1;
    mulTable.push_back({x, y, dy, d2y, 0});
    while(y != 0 && err > tolerance && dy2 != y*d2y && count < maxIt) {
        x1 = x - (y*dy)/(dy2-y*d2y);
        y = f(x1);
        dy = df(x1);
        d2y = d2f(x1);
        dy2 = pow(dy, 2);
        err = abs(x1 - x);
        x = x1;
        count++;
        mulTable.push_back({x, y, dy, d2y, (int)count});
    }
    ans.iterations = count;
    if(y == 0) {
        ans.isRoot = true;
        ans.root = x;
        return ans;
    }
    if(err <= tolerance) {
        ans.isAlmostRoot = true;
        ans.root = x;
        return ans;
    }
    if(dy == 0) {
        ans.badIn = true;
        return ans;
    }
    return ans;
    
}
