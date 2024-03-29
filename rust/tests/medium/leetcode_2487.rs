use rust::medium::leetcode_2487::ListNode;
use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn make_linked(mut nums: Vec<i32>) -> Option<Box<ListNode>> {
    if nums.is_empty() {
        return None;
    }

    let mut node = ListNode::new(nums.pop().unwrap_or(0));
    node.next = make_linked(nums);

    Some(Box::new(node))
}

fn output_linked(mut node: Option<Box<ListNode>>) -> Vec<i32> {
    let mut nums: Vec<i32> = Vec::new();
    while node.is_some() {
        if let Some(cur) = node {
            nums.push(cur.val);
            node = cur.next;
        }
    }

    nums
}

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, i32>, TargetData<Vec<i32>>>> {
    vec![
        DataCarrier::new(
            SourceData::new(vec![5,2,13,3,8], 0),
            TargetData::new(vec![13,8]),
        ),
        DataCarrier::new(
            SourceData::new(vec![1,1,1,1], 0),
            TargetData::new(vec![1, 1, 1, 1]),
        ),
    ]
}

#[cfg(test)]
mod test {
    use rust::medium::leetcode_2487::RemoveNodes;
    use crate::medium::leetcode_2487::{data_set, make_linked, output_linked};

    #[test]
    fn case_one() {
        for data in data_set() {
            let mut  value = data.source_data.value;
            value.reverse();
            let nodes = make_linked(value);

            assert_eq!(
                output_linked(RemoveNodes::case_one(nodes)),
                data.target_data.value
            );
        }
    }
}