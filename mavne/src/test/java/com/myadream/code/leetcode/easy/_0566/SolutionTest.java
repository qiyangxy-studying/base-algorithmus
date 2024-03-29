package com.myadream.code.leetcode.easy._0566;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> datasets = new ArrayList<>();
        Assist assist = new Assist(new int[][]{}, 0, 0);

        datasets.add(new DataSet(
                assist.copy(new int[][]{{1, 2}, {3, 4}},1, 4),
                new int[][]{{1, 2, 3, 4}}
        ));

        datasets.add(new DataSet(
                assist.copy(new int[][]{{1, 2}, {3, 4}}, 2, 4),
                new int[][]{{1, 2}, {3, 4}}
        ));

        return datasets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();

        impls.add(
                dataSet -> {
                    Assist assist = (Assist) dataSet.getSample();
                    assertArrayEquals((int[][]) dataSet.getTarget(), solution.matrixReshape(
                            assist.getSample(),
                            assist.getR(),
                            assist.getC()
                    ));

                    return true;
                }
        );

        return impls;
    }
}