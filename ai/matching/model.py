# Cài đặt thư viện (nếu chưa cài đặt, bạn cần chạy lệnh này trên terminal)
# pip install sentence-transformers

from sentence_transformers import SentenceTransformer, util

# Khởi tạo mô hình


def ranking_approriate_project(student_text, enterprise_texts, threshold=0.6):
    """

    hàm này trả về list các công việc theo thứ tự giảm dần của độ tương đồng cosine giữa văn bản của sinh viên và văn bản của doanh nghiệp 


    """
    # Tính embedding cho từng văn bản
    model = SentenceTransformer('matching_model')
    embedding_student = model.encode(student_text, convert_to_tensor=True)
    cosine_similarity_list = [] 
    for id, enterprise_text in enterprise_texts: 

        embedding_enterprise = model.encode(enterprise_text, convert_to_tensor=True) 
        
        cosine_sim = util.cos_sim(embedding_student, embedding_enterprise).item()
        
        print( enterprise_text, cosine_sim ) 
        
        # Lưu trữ kết quả nếu độ tương đồng vượt ngưỡng
        if cosine_sim >= threshold:
            cosine_similarity_list.append((id, cosine_sim))

    # Sắp xếp danh sách theo độ tương đồng giảm dần
    cosine_similarity_list.sort(key=lambda x: x[1], reverse=True)
    
    k = min(10, len(cosine_similarity_list))

    return cosine_similarity_list[k-1]
# Ví dụ: nhập văn bản của sinh viên và doanh nghiệp

def generate_score(student_text, enterprise_text):
    model = SentenceTransformer('all-MiniLM-L6-v2')
    embedding_student = model.encode(student_text, convert_to_tensor=True)
    embedding_enterprise = model.encode(enterprise_text, convert_to_tensor=True)

    return util.cos_sim(embedding_student, embedding_enterprise).item() * 100

if __name__ == "__main__":


    


    # nguyện vọng của ứng viên / sinh viên | Dạng text | 

    student_aspiration = "I aspire to develop groundbreaking AI technology solutions and participate in innovative projects."


    # Yêu cầu của danh nghiệp | Dạng lish các cặp [ id của project , yêu cầu của doanh nghiệp ]
    enterprise_requirements = [ [1, "We are seeking students who are passionate about technology and possess creative thinking to implement artificial intelligence projects."] , 
                            [2, "We are seeking students who are passionate about technology and possess creative thinking to implement website projects."] ] 


    # Kiểm tra khả năng hợp tác
    approriate_projects = ranking_approriate_project(student_aspiration, enterprise_requirements )

    print(approriate_projects)